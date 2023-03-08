package account

import (
	"fmt"
	"net/http"
	accountDB "rena-server-v2/internal/account/database"
	"rena-server-v2/internal/account/model"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/middleware"
	"rena-server-v2/internal/middleware/handler"
	"rena-server-v2/pkg/logging"
	"rena-server-v2/pkg/validate"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	accountDB accountDB.AccountDB
}

// signUp handles POST /v1/api/users
func (h *Handler) signUp(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		type RequestBody struct {
			User struct {
				Address  string `json:"address" binding:"required"`
				Username string `json:"username" binding:"required"`
			} `json:"user"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("account.handler.signUp failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.User, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid user request in body", details)
		}

		acc := model.Account{
			Address:  body.User.Address,
			Username: body.User.Username,
		}
		err := h.accountDB.Save(c.Request.Context(), &acc)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate email address", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewUserResponse(&acc))
	})
}

// currentUser handles GET /v1/api/user/me
func (h *Handler) currentUser(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		currentUser := MustCurrentUser(c)
		find, err := h.accountDB.FindByAddress(c.Request.Context(), currentUser.Address)
		if err != nil {
			if database.IsRecordNotFoundErr(err) {
				return handler.NewErrorResponse(http.StatusNotFound, handler.NotFoundEntity, "not found current user", nil)
			}
			return &handler.Response{Err: err}
		}
		return handler.NewSuccessResponse(http.StatusOK, NewUserResponse(find))
	})
}

// users handles GET /v1/api/users
func (h *Handler) users(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type QueryParameter struct {
			Address string `form:"address" binding:"omitempty,len=42"`
		}
		var query QueryParameter
		if err := c.ShouldBindQuery(&query); err != nil {
			logger.Errorw("account.handler.users failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&query, "form", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid item request in query", details)
		}

		fmt.Println(query)
		criteria := accountDB.IterateUserCriteria{
			Address: query.Address,
		}
		users, total, err := h.accountDB.FindUsers(c.Request.Context(), criteria)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewUsersResponse(users, total))
	})
}

// user handles GET /v1/api/users
func (h *Handler) user(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type QueryParameter struct {
			Address string `form:"address" binding:"omitempty,len=42"`
		}
		var query QueryParameter
		if err := c.ShouldBindQuery(&query); err != nil {
			logger.Errorw("account.handler.users failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&query, "form", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid item request in query", details)
		}

		fmt.Println(query)
		user, err := h.accountDB.FindByAddress(c.Request.Context(), query.Address)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewUserResponse(user))
	})
}

// update handles PUT /v1/api/user
func (h *Handler) update(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		currentUser := MustCurrentUser(c)
		type RequestBody struct {
			User struct {
				Address  string `json:"address" binding:"required"`
				Username string `json:"username" binding:"omitempty"`
				Bio      string `json:"bio"`
				Image    string `json:"image"`
			} `json:"user"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("account.handler.update failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.User, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid user request in body", details)
		}

		acc, err := h.accountDB.FindByAddress(c.Request.Context(), currentUser.Address)
		if err != nil {
			if database.IsRecordNotFoundErr(err) {
				return handler.NewErrorResponse(http.StatusNotFound, handler.NotFoundEntity, "not found account", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}

		if body.User.Address != "" {
			acc.Address = body.User.Address
		}
		if body.User.Username != "" {
			acc.Username = body.User.Username
		}
		if body.User.Bio != "" {
			acc.Bio = body.User.Bio
		}
		if body.User.Image != "" {
			acc.Image = body.User.Image
		}
		err = h.accountDB.Update(c.Request.Context(), currentUser.Address, acc)
		if err != nil {
			if database.IsRecordNotFoundErr(err) {
				logger.Errorw("account.handler.update failed to update user because not found user", "err", err)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewUserResponse(acc))
	})
}

// RouteV1 routes user api given config and gin.Engine
func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("v1/api")
	timeout := time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second
	v1.Use(middleware.RequestIDMiddleware(), middleware.TimeoutMiddleware(timeout))
	// anonymous
	v1.Use()
	{
		v1.POST("users/login", auth.LoginHandler)
		v1.POST("users", h.signUp)
		v1.GET("users", h.users)
		v1.GET("user", h.user)
		v1.GET("user/me", h.currentUser)
		v1.PUT("user", h.update)
	}
	// auth required
	v1.Use(auth.MiddlewareFunc())
	{
	}
}

func NewHandler(accountDB accountDB.AccountDB) *Handler {
	return &Handler{
		accountDB: accountDB,
	}
}
