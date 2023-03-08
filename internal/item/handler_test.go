package item

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"rena-server-v2/internal/account"
	accountDBMock "rena-server-v2/internal/account/database/mocks"
	accountModel "rena-server-v2/internal/account/model"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/item/database"
	itemDBMock "rena-server-v2/internal/item/database/mocks"
	"rena-server-v2/internal/item/model"
	"rena-server-v2/pkg/logging"
	"strings"
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

	dItem = model.Item{
		ID:        1,
		Slug:      "how-to-train-your-dragon",
		Title:     "How to train your dragon",
		Body:      "You have to believe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Author:    dUser,
		AuthorID:  dUser.ID,
		Tags: []*model.Tag{
			{
				ID:        3,
				Name:      "reactjs",
				CreatedAt: time.Now(),
			}, {
				ID:        2,
				Name:      "angularjs",
				CreatedAt: time.Now(),
			}, {
				ID:        1,
				Name:      "dragons",
				CreatedAt: time.Now(),
			},
		},
	}
	dItemTags = []string{"reactjs", "angularjs", "dragons"}
)

type HandlerSuite struct {
	suite.Suite
	r         *gin.Engine
	handler   *Handler
	db        *itemDBMock.ItemDB
	accountDB *accountDBMock.AccountDB
}

func (s *HandlerSuite) SetupSuite() {
	logging.SetLevel(zapcore.FatalLevel)
}

func (s *HandlerSuite) SetupTest() {
	cfg, err := config.Load("")
	s.NoError(err)

	s.db = &itemDBMock.ItemDB{}
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

func (s *HandlerSuite) TestSaveItem() {
	// given
	s.db.On("SaveItem", mock.Anything, mock.Anything).Return(nil)

	// when
	requestBody := map[string]interface{}{
		"item": map[string]interface{}{
			"title":   dItem.Title,
			"body":    dItem.Body,
			"tagList": dItemTags,
		},
	}
	b, _ := json.Marshal(&requestBody)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/items", bytes.NewBuffer(b))
	req.Header.Add("Authorization", "Bearer "+s.getBearerToken())

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	itemMatcher := itemMatcher(dItem.Title, dItem.Body, dItemTags, &dUser)
	s.db.AssertCalled(s.T(), "SaveItem", mock.Anything, mock.MatchedBy(itemMatcher))
	// 2) status code
	s.Equal(http.StatusCreated, res.Code)
	// 3) response
	jsonVal := res.Body.String()
	s.assertItemResponse(&dItem, gjson.Parse(jsonVal).Get("item"))
}

func (s *HandlerSuite) TestItemBySlug() {
	// given
	s.db.On("FindItemBySlug", mock.Anything, dItem.Slug).Return(&dItem, nil)

	// when
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/items/"+dItem.Slug, nil)

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	s.db.AssertCalled(s.T(), "FindItemBySlug", mock.Anything, dItem.Slug)
	// 2) status code
	s.Equal(http.StatusOK, res.Code)
	// 3) body
	jsonVal := res.Body.String()
	s.assertItemResponse(&dItem, gjson.Parse(jsonVal).Get("item"))
}

func (s *HandlerSuite) TestItems() {
	criteria := database.IterateItemCriteria{
		Tags:   []string{dItemTags[0]},
		Author: dItem.Author.Username,
		Offset: 0,
		Limit:  5,
	}
	s.db.On("FindItems", mock.Anything, criteria).Return([]*model.Item{&dItem}, int64(1), nil)

	// when
	url := fmt.Sprintf("/v1/api/items?tag=%s&author=%s&offset=%d&limit=%d", criteria.Tags[0],
		criteria.Author, criteria.Offset, criteria.Limit)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)

	s.r.ServeHTTP(res, req)

	// then
	// 1) method called
	s.db.AssertCalled(s.T(), "FindItems", mock.Anything, criteria)
	// 2) status code
	s.Equal(http.StatusOK, res.Code)
	// 3) body
	jsonVal := res.Body.String()
	result := gjson.Parse(jsonVal)
	s.Equal(int64(1), result.Get("itemsCount").Int())

	itemsResult := result.Get("items").Array()
	s.Equal(1, len(itemsResult))
	s.assertItemResponse(&Item, itemsResult[0])
}

func (s *HandlerSuite) TestDeleteItem() {
	// given
	s.db.On("RunInTx", mock.Anything, mock.Anything).Return(nil)

	// when
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/v1/api/items/"+dItem.Slug, nil)
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

func (s *HandlerSuite) assertItemResponse(item *model.Item, result gjson.Result) {
	s.Equal(slug.Make(item.Title), result.Get("slug").String())
	s.Equal(item.Title, result.Get("title").String())
	s.Equal(item.Body, result.Get("body").String())

	var tagVals []string
	for _, tag := range item.Tags {
		tagVals = append(tagVals, tag.Name)
	}
	for _, result := range result.Get("tagList").Array() {
		s.Contains(tagVals, result.String())
	}
	s.True(result.Get("createdAt").Exists())
	s.True(result.Get("updatedAt").Exists())
	s.Equal(item.Author.Username, result.Get("author.username").String())
	s.Equal(item.Author.Bio, result.Get("author.bio").String())
	s.Equal(item.Author.Image, result.Get("author.image").String())
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

func itemMatcher(title, body string, tags []string, author *accountModel.Account) func(a *model.Item) bool {
	return func(a *model.Item) bool {
		if a.Slug != slug.Make(title) || a.Title != title || a.Body != body {
			return false
		}
		if len(tags) != len(a.Tags) {
			return false
		}
		tagVals := strings.Join(tags, " ")
		for _, tag := range a.Tags {
			if !strings.Contains(tagVals, tag.Name) {
				return false
			}
		}
		if a.Author.Username != author.Username || a.Author.Bio != author.Bio || a.Author.Image != author.Image {
			return false
		}
		return true
	}
}
