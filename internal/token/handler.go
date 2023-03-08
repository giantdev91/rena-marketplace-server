package token

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	"rena-server-v2/internal/middleware"
	"rena-server-v2/internal/middleware/handler"
	tokenDB "rena-server-v2/internal/token/database"
	"rena-server-v2/internal/token/model"
	"rena-server-v2/pkg/logging"
	"rena-server-v2/pkg/validate"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"

	tokenContract "rena-server-v2/internal/contract/token"
)

type Handler struct {
	tokenDB tokenDB.TokenDB
}

// tokens handles GET /v1/api/tokens
func (h *Handler) tokens(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type QueryParameter struct {
			Offset uint `form:"offset,default=0" binding:"numeric"`
			Limit  uint `form:"limit,default=6" binding:"numeric"`
		}
		var query QueryParameter
		if err := c.ShouldBindQuery(&query); err != nil {
			logger.Errorw("token.handler.tokens failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&query, "form", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid token request in query", details)
		}

		fmt.Println(query)
		criteria := tokenDB.IterateTokenCriteria{
			Offset: query.Offset,
			Limit:  query.Limit,
		}
		tokens, total, err := h.tokenDB.FindTokens(c.Request.Context(), criteria)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewTokensResponse(tokens, total))
	})
}

// saveToken handles POST /v1/api/tokens
func (h *Handler) saveToken(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Token struct {
				Token string `json:"token" binding:"omitempty,len=42"`
			} `json:"token"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("token.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Token, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid token request in body", details)
		}

		cfg, err := config.Load("")
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}

		client, err := ethclient.Dial(cfg.CryptoConfig.Network)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}

		privateKey, err := crypto.HexToECDSA(cfg.CryptoConfig.PrivateKey)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			return handler.NewInternalErrorResponse(err)
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}

		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0) // in wei

		tokenAddress := common.HexToAddress(cfg.CryptoConfig.TokenRegistry)
		tokenInstance, err := tokenContract.NewFantomTokenRegistry(tokenAddress, client)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}

		address := common.HexToAddress(body.Token.Token)
		gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
		tx, err := tokenInstance.Add(&bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasPrice: gasPrice,
		}, address)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}
		fmt.Println(tx)

		// save token
		token := model.Token{
			Token: body.Token.Token,
		}
		err = h.tokenDB.SaveToken(c.Request.Context(), &token)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate token", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewTokenResponse(&token))
	})
}

// deleteToken handles DELETE /v1/api/token/:address
func (h *Handler) deleteToken(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)
		// bind
		type RequestUri struct {
			Address string `uri:"address" binding:"required"`
		}
		var uri RequestUri
		if err := c.ShouldBindUri(&uri); err != nil {
			logger.Errorw("article.handler.deleteToken failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&uri, "uri", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid article request in uri", details)
		}

		// delete token with in transaction
		err := h.tokenDB.RunInTx(c.Request.Context(), func(ctx context.Context) error {
			// delete a token
			if err := h.tokenDB.DeleteToken(ctx, uri.Address); err != nil {
				return err
			}
			logger.Debugw("token.handler.deleteToken success to delete a token")
			return nil
		})
		if err != nil {
			logger.Errorw("token.handler.deleteToken failed to delete a token", "err", err)
			if database.IsRecordNotFoundErr(errors.Cause(err)) {
				return handler.NewErrorResponse(http.StatusNotFound, handler.NotFoundEntity, "not found token", nil)
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

	tokenV1 := v1.Group("tokens")
	// anonymous
	tokenV1.Use()
	{
		tokenV1.GET("", h.tokens)
		tokenV1.POST("", h.saveToken)
		tokenV1.DELETE(":address", h.deleteToken)
	}
}

func NewHandler(tokenDB tokenDB.TokenDB) *Handler {
	return &Handler{
		tokenDB: tokenDB,
	}
}
