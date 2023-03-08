package collection

import (
	"context"
	"net/http"
	collectionDB "rena-server-v2/internal/collection/database"
	"rena-server-v2/internal/collection/model"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/middleware"
	"rena-server-v2/internal/middleware/handler"
	"rena-server-v2/pkg/logging"
	"rena-server-v2/pkg/validate"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
)

type Handler struct {
	collectionDB collectionDB.CollectionDB
}

// saveCollection handles POST /v1/api/collections
func (h *Handler) saveCollection(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		// bind
		type RequestBody struct {
			Collection struct {
				ContractAddress string `json:"contractAddress" binding:"required"`
				Name            string `json:"name" binding:"required"`
				Description     string `json:"description" binding:"omitempty"`
				ImageURL        string `json:"imageURL" binding:"omitempty"`
				Royalty         int    `json:"royalty" binding:"omitempty"`
				FeeRecipient    string `json:"feeRecipient" binding:"required"`
				Category        string `json:"category" binding:"omitempty"`
				Website         string `json:"website" binding:"omitempty"`
				Discord         string `json:"discord" binding:"omitempty"`
				TwitterUrl      string `json:"twitterUrl" binding:"omitempty"`
				InstagramUrl    string `json:"instagramUrl" binding:"omitempty"`
				MediumUrl       string `json:"mediumUrl" binding:"omitempty"`
				TelegramUrl     string `json:"telegramUrl" binding:"omitempty"`
				ContactEmail    string `json:"contactEmail" binding:"required,email"`
			} `json:"collection"`
		}
		var body RequestBody

		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("collection.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Collection, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid collection request in body", details)
		}

		// save collection
		collection := model.Collection{
			Slug:            slug.Make(body.Collection.Name),
			ContractAddress: body.Collection.ContractAddress,
			Name:            body.Collection.Name,
			Description:     body.Collection.Description,
			ImageURL:        body.Collection.ImageURL,
			Royalty:         body.Collection.Royalty,
			FeeRecipient:    body.Collection.FeeRecipient,
			Category:        body.Collection.Category,
			Website:         body.Collection.Website,
			Discord:         body.Collection.Discord,
			TwitterUrl:      body.Collection.TwitterUrl,
			InstagramUrl:    body.Collection.InstagramUrl,
			MediumUrl:       body.Collection.MediumUrl,
			TelegramUrl:     body.Collection.TelegramUrl,
			ContactEmail:    body.Collection.ContactEmail,
		}
		err := h.collectionDB.SaveCollection(c.Request.Context(), &collection)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate collection name", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewCollectionResponse(&collection))
	})
}

// collectionBySlug handles GET /v1/api/collections/:name
func (h *Handler) collectionBySlug(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		// bind
		type RequestUri struct {
			Slug string `uri:"slug"`
		}
		var uri RequestUri
		if err := c.ShouldBindUri(&uri); err != nil {
			logger.Errorw("collection.handler.collectionBySlug failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&uri, "uri", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid collection request in uri", details)
		}

		// find
		collection, err := h.collectionDB.FindCollectionBySlug(c.Request.Context(), uri.Slug)
		if err != nil {
			if database.IsRecordNotFoundErr(err) {
				return handler.NewErrorResponse(http.StatusNotFound, handler.NotFoundEntity, "not found collection", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewCollectionResponse(collection))
	})
}

// collections handles GET /v1/api/collections
func (h *Handler) collections(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		type QueryParameter struct {
			Slug   string `form:"slug" binding:"omitempty"`
			Limit  string `form:"limit,default=5" binding:"numeric"`
			Offset string `form:"offset,default=0" binding:"numeric"`
		}
		var query QueryParameter
		if err := c.ShouldBindQuery(&query); err != nil {
			logger.Errorw("collection.handler.collections failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&query, "form", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid collection request in query", details)
		}

		limit, err := strconv.ParseUint(query.Limit, 10, 64)
		if err != nil {
			limit = 5
		}
		offset, err := strconv.ParseUint(query.Offset, 10, 64)
		if err != nil {
			offset = 0
		}
		criteria := collectionDB.IterateCollectionCriteria{
			Slug:   query.Slug,
			Offset: uint(offset),
			Limit:  uint(limit),
		}
		collections, total, err := h.collectionDB.FindCollections(c.Request.Context(), criteria)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewCollectionsResponse(collections, total))
	})
}

// updateCollection handles POST /v1/api/collections/update
func (h *Handler) updateCollection(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Collection struct {
				ContractAddress string `json:"contractAddress" binding:"required,len=42"`
				Name            string `json:"name" binding:"required"`
				ImageURL        string `json:"imageURL" binding:"required"`
				Description     string `json:"description" binding:"required"`
				Royalty         int    `json:"royalty" binding:"omitempty"`
				FeeRecipient    string `json:"feeRecipient" binding:"required"`
				Category        string `json:"category" binding:"omitempty"`
				WebsiteUrl      string `json:"websiteUrl" binding:"omitempty"`
				DiscordUrl      string `json:"discordUrl" binding:"omitempty"`
				TwitterUrl      string `json:"twitterUrl" binding:"omitempty"`
				InstagramUrl    string `json:"instagramUrl" binding:"omitempty"`
				MediumUrl       string `json:"mediumUrl" binding:"omitempty"`
				TelegramUrl     string `json:"telegramUrl" binding:"omitempty"`
				ContactEmail    string `json:"contactEmail" binding:"omitempty"`
			} `json:"collection"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("collection.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Collection, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid collection request in body", details)
		}

		collection := model.Collection{
			Slug:            slug.Make(body.Collection.Name),
			ContractAddress: body.Collection.ContractAddress,
			Name:            body.Collection.Name,
			Description:     body.Collection.Description,
			ImageURL:        body.Collection.ImageURL,
			Royalty:         body.Collection.Royalty,
			FeeRecipient:    body.Collection.FeeRecipient,
			Category:        body.Collection.Category,
			Website:         body.Collection.WebsiteUrl,
			Discord:         body.Collection.DiscordUrl,
			TwitterUrl:      body.Collection.TwitterUrl,
			InstagramUrl:    body.Collection.InstagramUrl,
			MediumUrl:       body.Collection.MediumUrl,
			TelegramUrl:     body.Collection.TelegramUrl,
			ContactEmail:    body.Collection.ContactEmail,
		}
		// update item
		err := h.collectionDB.UpdateCollection(c.Request.Context(), &collection)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewCollectionResponse(&collection))
	})
}

// deleteCollection handles DELETE /v1/api/collections/:slug
func (h *Handler) deleteCollection(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		// bind
		type RequestUri struct {
			Slug string `uri:"slug" binding:"required"`
		}
		var uri RequestUri
		if err := c.ShouldBindUri(&uri); err != nil {
			logger.Errorw("collection.handler.deleteCollection failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&uri, "uri", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid collection request in uri", details)
		}
		// delete collection with in transaction
		err := h.collectionDB.RunInTx(c.Request.Context(), func(ctx context.Context) error {
			// delete a collection
			if err := h.collectionDB.DeleteCollectionBySlug(ctx, uri.Slug); err != nil {
				return err
			}

			logger.Debugw("collection.handler.deleteCollection success to delete a collection")
			return nil
		})
		if err != nil {
			logger.Errorw("collection.handler.deleteCollection failed to delete a collection", "err", err)
			if database.IsRecordNotFoundErr(errors.Cause(err)) {
				return handler.NewErrorResponse(http.StatusNotFound, handler.NotFoundEntity, "not found collection", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, nil)
	})
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("v1/api")
	timeout := time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second
	v1.Use(middleware.RequestIDMiddleware(), middleware.TimeoutMiddleware(timeout))

	collectionV1 := v1.Group("collections")
	// anonymous
	collectionV1.Use()
	{
		collectionV1.GET(":slug", h.collectionBySlug)
		collectionV1.GET("", h.collections)
		collectionV1.POST("", h.saveCollection)
		collectionV1.PUT("", h.updateCollection)
		collectionV1.DELETE(":slug", h.deleteCollection)
	}

	// auth required
	collectionV1.Use(auth.MiddlewareFunc())
	{
	}
}

func NewHandler(collectionDB collectionDB.CollectionDB) *Handler {
	return &Handler{
		collectionDB: collectionDB,
	}
}
