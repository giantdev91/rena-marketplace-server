package account

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"rena-server-v2/internal/account/database/mocks"
	"rena-server-v2/internal/account/model"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	"rena-server-v2/pkg/logging"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"
	"go.uber.org/zap/zapcore"
)

type HandlerSuite struct {
	suite.Suite
	r       *gin.Engine
	handler *Handler
	db      *mocks.AccountDB
}

func (s *HandlerSuite) SetupSuite() {
	logging.SetLevel(zapcore.FatalLevel)
}

func (s *HandlerSuite) SetupTest() {
	cfg, err := config.Load("")
	s.NoError(err)

	s.db = &mocks.AccountDB{}
	s.handler = NewHandler(s.db)

	jwtMiddleware, err := NewAuthMiddleware(cfg, s.db)
	s.NoError(err)

	gin.SetMode(gin.TestMode)
	s.r = gin.Default()

	RouteV1(cfg, s.handler, s.r, jwtMiddleware)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

func (s *HandlerSuite) TestRegister() {
	// given
	username := "zaccoding"
	email := "zaccoding@gmail.com"

	s.db.On("Save", mock.Anything).Return(nil)

	// when
	body := map[string]interface{}{
		"user": map[string]interface{}{
			"username": username,
			"email":    email,
		},
	}
	b, _ := json.Marshal(body)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/users", bytes.NewBuffer(b))

	s.r.ServeHTTP(res, req)

	// then
	s.db.AssertCalled(s.T(), "Save", mock.Anything)
	s.Equal(http.StatusCreated, res.Code)
	expected := `
		{
		  "user": {
			"username": "zaccoding",
			"email": "zaccoding@gmail.com",
			"bio": "",
			"image": ""
		  }
		}`
	s.JSONEq(expected, res.Body.String())
}

func (s *HandlerSuite) TestRegister_BadRequest() {
	username := "user1"
	email := "user@email.com"
	cases := []struct {
		Address  string
		Username string
		Email    string
		Expected string
	}{
		{
			Expected: `
			{
			  "code": "InvalidBodyValue",
			  "message": "[InvalidBodyValue] invalid user request in body",
			  "errors": [
				{
				  "field": "username",
				  "value": "",
				  "message": "required username"
				},
				{
				  "field": "email",
				  "value": "",
				  "message": "required email"
				},
				{
				  "field": "password",
				  "value": "",
				  "message": "required password"
				}
			  ]
			}`,
		},
		// username
		{
			Email: email,
			Expected: `
			{
			  "code": "InvalidBodyValue",
			  "message": "[InvalidBodyValue] invalid user request in body",
			  "errors": [
				{
				  "field": "username",
				  "value": "",
				  "message": "required username"
				}
			  ]
			}`,
		},
		// email
		{
			Username: username,
			Expected: `
			{
			  "code": "InvalidBodyValue",
			  "message": "[InvalidBodyValue] invalid user request in body",
			  "errors": [
				{
				  "field": "email",
				  "value": "",
				  "message": "required email"
				}
			  ]
			}`,
		}, {
			Username: username,
			Email:    "invalid-email-format",
			Expected: `
			{
			  "code": "InvalidBodyValue",
			  "message": "[InvalidBodyValue] invalid user request in body",
			  "errors": [
				{
				  "field": "email",
				  "value": "invalid-email-format",
				  "message": "required email format"
				}
			  ]
			}`,
		},
		// password
		{
			Username: username,
			Email:    email,
			Expected: `
			{
			  "code": "InvalidBodyValue",
			  "errors": [
				{
				  "field": "password",
				  "value": "",
				  "message": "required password"
				}
			  ],
			  "message": "[InvalidBodyValue] invalid user request in body"
			}`,
		}, {
			Username: username,
			Email:    email,
			Expected: `
			{
			  "code": "InvalidBodyValue",
			  "errors": [
				{
				  "field": "password",
				  "value": "1",
				  "message": "password required at least 5 length"
				}
			  ],
			  "message": "[InvalidBodyValue] invalid user request in body"
			}`,
		},
	}

	for _, tc := range cases {
		userReq := map[string]interface{}{}
		if tc.Username != "" {
			userReq["username"] = tc.Username
		}
		if tc.Email != "" {
			userReq["email"] = tc.Email
		}

		b, _ := json.Marshal(map[string]interface{}{
			"user": userReq,
		})
		res := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/api/users", bytes.NewBuffer(b))

		s.r.ServeHTTP(res, req)

		// then
		s.db.AssertNotCalled(s.T(), "Save", mock.Anything, mock.Anything)
		s.Equal(http.StatusBadRequest, res.Code)
		s.JSONEq(tc.Expected, res.Body.String())
	}
}

func (s *HandlerSuite) TestRegister_FailIfDuplicateError() {
	// given
	username := "zaccoding"
	email := "zaccoding@gmail.com"
	s.db.On("Save", mock.Anything).Return(database.ErrKeyConflict)

	// when
	body := map[string]interface{}{
		"user": map[string]interface{}{
			"username": username,
			"email":    email,
		},
	}
	b, _ := json.Marshal(body)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/users", bytes.NewBuffer(b))

	s.r.ServeHTTP(res, req)

	// then
	s.db.AssertCalled(s.T(), "Save", mock.Anything)
	s.Equal(http.StatusConflict, res.Code)
	expected := `
	{
	  "code": "DuplicateEntry",
	  "message": "[DuplicateEntry] duplicate email address"
	}`
	s.JSONEq(expected, res.Body.String())
}

func (s *HandlerSuite) TestCurrentUser() {
	// given
	address := "0x990f5Df13b6894FbA311Af4069bcD425eaf954ed"
	acc := model.Account{
		ID:        1,
		Username:  "user1",
		Email:     "user1@gmail.com",
		Bio:       "user1 bio",
		Image:     "user1 image",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	token := s.getBearerToken(&acc, address)

	// when
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/user/me", nil)
	req.Header.Add("Authorization", "Bearer "+token)

	s.r.ServeHTTP(res, req)

	// then
	s.db.AssertCalled(s.T(), "FindByAddress", mock.Anything, acc.Address)
	s.Equal(http.StatusOK, res.Code)
	expected := `
	{
	  "user": {
		"username": "user1",
		"email": "user1@gmail.com",
		"bio": "user1 bio",
		"image": "user1 image"
	  }
	}`
	s.JSONEq(expected, res.Body.String())
}

func (s *HandlerSuite) TestUpdate() {
	// given
	address := "0x990f5Df13b6894FbA311Af4069bcD425eaf954ed"
	acc := model.Account{
		ID:        1,
		Address:   "0x990f5Df13b6894FbA311Af4069bcD425eaf954ed",
		Username:  "user1",
		Email:     "user1@gmail.com",
		Bio:       "user1 bio",
		Image:     "user1 image",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	token := s.getBearerToken(&acc, address)
	s.db.On("Update", mock.Anything, acc.Address, mock.Anything).Return(nil)

	// when
	updateRequest := map[string]interface{}{
		"user": map[string]interface{}{
			"username": "updated-user1",
			"bio":      "updated-bio",
			"image":    "updated-image",
		},
	}
	b, _ := json.Marshal(updateRequest)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/v1/api/user", bytes.NewBuffer(b))
	req.Header.Add("Authorization", "Bearer "+token)

	s.r.ServeHTTP(res, req)

	// then
	s.db.AssertCalled(s.T(), "Update", mock.Anything, acc.Email, mock.MatchedBy(func(a *model.Account) bool {
		return a.Email == acc.Email && a.Username == "updated-user1" && a.Bio == "updated-bio" && a.Image == "updated-image"
	}))
	s.Equal(http.StatusOK, res.Code)
	expected := `
	{
	  "user": {
			"username": "updated-user1",
			"email": "user1@gmail.com",
			"bio": "updated-bio",
			"image": "updated-image"
	  }
	}`
	s.JSONEq(expected, res.Body.String())
}

func (s *HandlerSuite) getBearerToken(acc *model.Account, address string) string {
	s.db.On("FindByAddress", mock.Anything, acc.Address).Return(acc, nil)
	body := map[string]interface{}{
		"user": map[string]interface{}{
			"address": acc.Address,
		},
	}
	b, _ := json.Marshal(body)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/api/users/login", bytes.NewBuffer(b))
	s.r.ServeHTTP(res, req)

	return gjson.Get(res.Body.String(), "token").String()
}
