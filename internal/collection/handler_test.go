package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rena-server-v2/internal/account"
	accountDBMock "rena-server-v2/internal/account/database/mocks"
	accountModel "rena-server-v2/internal/account/model"
	"rena-server-v2/internal/collection/database"
	"rena-server-v2/internal/collection/model"
	"rena-server-v2/internal/config"
	"rena-server-v2/pkg/logging"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
	"go.uber.org/zap/zapcore"
)

var (
	dUser = accountModel.Account{
		ID:       1,
		Username: "user1",
		Email:    "user1@gmail.com",
		Password: "$2a$10$lsYsLv8nGPM0.R.ft4sgpe3OP7..KL3ZJqqhSVCKTEnSCMUztoUcW",
		Bio:      "I am working!",
	}
	dUserRawPass = "user1"

	dCollection = model.Collection{
		ID:              1,
		ContractAddress: "how-to-create-your-collection",
		Name:            "How to create your collection",
		Description:     "You have to believe",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
)

type HandlerSuite struct {
	suite.Suite
	r         *gin.Engine
	handler   *Handler
	db        *collectionDBMock.CollectionDB
	accountDB *accountDBMock.AccountDB
}

func (s *HandlerSuite) SetupSuite() {
	logging.SetLevel(zapcore.FatalLevel)
}

func (s *HandlerSuite) SetupTest() {
	cfg, err := config.Load("")
	s.NoError(err)

	s.db = &collectionDBMock.CollectionDB{}
	s.handler = NewHandler(s.db)
	s.accountDB = &accountDBMock.AccountDB{}
	s.accountDB.On("FindByEmail", mock.Anything, mock.MatchedBy(func(email string) bool {
		return email == dUser.Email
	})).Return(&dUser, nil)

	jwtMiddleware, err := account.NewAuthMiddleware(cfg, s.accountDB)
	s.NoError(err)

	gin.SetMode(gin.TestMode)
	s.r = gin.Default()

	RouteV1(cfg, s.handler, s.r, jwtMiddleware)

	accountHandler := account.NewHandler(s.accountDB)
	account.RouteV1(cfg, accountHandler, s.r, jwtMiddleware)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

// TODO : test failures

func (s *HandlerSuite) TestSaveCollection() {
	// given
	s.db.On("SaveCollection", mock.Anything, mock.Anything).Return(nil)

	// when
	requestBody := map[string]interface{}{
		"collection": map[string]interface{}{
			"contractAddress": dCollection.ContractAddress,
			"name":            dCollection.Name,
			"description":     dCollection.Description,
		},
	}
	b, _ := json.Marshal(&requestBody)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/collections", bytes.NewBuffer(b))
	req.Header.Add("Authorization", "Bearer "+s.getBearerToken())

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	collectionMatcher := collectionMatcher(dCollection.ContractAddress, dCollection.Name, &dUser)
	s.db.AssertCalled(s.T(), "SaveCollection", mock.Anything, mock.MatchedBy(collectionMatcher))
	// 2) status code
	s.Equal(http.StatusCreated, res.Code)
	// 3) response
	jsonVal := res.Body.String()
	s.assertCollectionResponse(&dCollection, gjson.Parse(jsonVal).Get("collection"))
}

func (s *HandlerSuite) TestCollectionBySlug() {
	// given
	s.db.On("FindCollectionBySlug", mock.Anything, dCollection.Name).Return(&dCollection, nil)

	// when
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/collections/"+dCollection.Name, nil)

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	s.db.AssertCalled(s.T(), "FindCollectionBySlug", mock.Anything, dCollection.Name)
	// 2) status code
	s.Equal(http.StatusOK, res.Code)
	// 3) body
	jsonVal := res.Body.String()
	s.assertCollectionResponse(&dCollection, gjson.Parse(jsonVal).Get("collection"))
}

func (s *HandlerSuite) TestCollections() {
	criteria := database.IterateCollectionCriteria{
		Slug:   dCollection.Slug,
		Offset: 0,
		Limit:  5,
	}
	s.db.On("FindCollections", mock.Anything, criteria).Return([]*model.Collection{&dCollection}, int64(1), nil)

	// when
	url := fmt.Sprintf("/v1/api/collections?slug=%s&offset=%d&limit=%d", criteria.Slug,
		criteria.Offset, criteria.Limit)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	s.db.AssertCalled(s.T(), "FindCollections", mock.Anything, criteria)
	// 2) status code
	s.Equal(http.StatusOK, res.Code)
	// 3) body
	jsonVal := res.Body.String()
	result := gjson.Parse(jsonVal)
	s.Equal(int64(1), result.Get("collectionsCount").Int())

	collectionsResult := result.Get("collections").Array()
	s.Equal(1, len(collectionsResult))
	s.assertCollectionResponse(&dCollection, collectionsResult[0])
}

func (s *HandlerSuite) TestDeleteCollection() {
	// given
	s.db.On("RunInTx", mock.Anything, mock.Anything).Return(nil)

	// when
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/api/collections/"+dCollection.Name, nil)
	req.Header.Add("Authorization", "Bearer "+s.getBearerToken())

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	s.db.AssertCalled(s.T(), "RunInTx", mock.Anything, mock.Anything)
	// 2) status code
	s.Equal(http.StatusOK, res.Code)
	// 3) body
	s.Empty(res.Body.Bytes())
}

func (s *HandlerSuite) assertCollectionResponse(collection *model.Collection, result gjson.Result) {
	s.Equal(slug.Make(collection.Name), result.Get("slug").String())
	s.Equal(collection.Name, result.Get("name").String())
	s.Equal(collection.Description, result.Get("description").String())

	s.True(result.Get("createdAt").Exists())
	s.True(result.Get("updatedAt").Exists())
	s.Equal(collection.Name, result.Get("name").String())
	s.Equal(collection.Description, result.Get("description").String())
	s.Equal(collection.ImageURL, result.Get("image_url").String())
}

func (s *HandlerSuite) getBearerToken() string {
	body := map[string]interface{}{
		"user": map[string]interface{}{
			"email":    dUser.Email,
			"password": dUserRawPass,
		},
	}
	b, _ := json.Marshal(body)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/users/login", bytes.NewBuffer(b))
	s.r.ServeHTTP(res, req)

	s.Equal(http.StatusOK, res.Code)
	return gjson.Get(res.Body.String(), "token").String()
}

func collectionMatcher(name, description string) func(a *model.Collection) bool {
	return func(a *model.Collection) bool {
		if a.Slug != slug.Make(name) || a.Name != name || a.Description != description {
			return false
		}
		return true
	}
}
