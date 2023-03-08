package marketplace

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	marketplaceDB "rena-server-v2/internal/marketplace/database"
	"rena-server-v2/internal/marketplace/model"
	"rena-server-v2/internal/middleware"
	"rena-server-v2/internal/middleware/handler"
	"rena-server-v2/pkg/logging"
	"rena-server-v2/pkg/validate"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	bundleMarketContract "rena-server-v2/internal/contract/bundlemarketplace"
	marketContract "rena-server-v2/internal/contract/marketplace"
)

type Handler struct {
	marketplaceDB marketplaceDB.MarketplaceDB
}

// saveMarketplace handles POST /v1/api/marketplaces/marketplace
func (h *Handler) saveMarketplace(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Marketplace struct {
				AddressRegistry string `json:"addressRegistry" binding:"omitempty,len=42"`
				PlatformFee     string `json:"platformFee" binding:"omitempty"`
				FeeRecipient    string `json:"feeRecipient" binding:"omitempty,len=42"`
			} `json:"marketplace"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("marketplace.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Marketplace, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid marketplace request in body", details)
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

		marketAddress := common.HexToAddress(cfg.CryptoConfig.Marketplace)
		marketInstance, err := marketContract.NewFantomMarketplace(marketAddress, client)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}

		if body.Marketplace.AddressRegistry != "" {
			address := common.HexToAddress(body.Marketplace.AddressRegistry)
			gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
			tx, err := marketInstance.UpdateAddressRegistry(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		if body.Marketplace.FeeRecipient != "" {
			address := common.HexToAddress(body.Marketplace.FeeRecipient)
			gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
			tx, err := marketInstance.UpdatePlatformFeeRecipient(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		if body.Marketplace.PlatformFee != "" {
			fee, err := strconv.ParseUint(body.Marketplace.PlatformFee, 10, 16)
			gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
			tx, err := marketInstance.UpdatePlatformFee(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, uint16(fee))
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		// save marketplace
		marketplace := model.Marketplace{
			AddressRegistry: body.Marketplace.AddressRegistry,
			PlatformFee:     body.Marketplace.PlatformFee,
			FeeRecipient:    body.Marketplace.FeeRecipient,
		}
		err = h.marketplaceDB.SaveMarketplace(c.Request.Context(), &marketplace)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate collection name", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewMarketplaceResponse(&marketplace))
	})
}

// saveBundleMarketplace handles POST /v1/api/marketplaces/bundlemarketplace
func (h *Handler) saveBundleMarketplace(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Marketplace struct {
				AddressRegistry string `json:"addressRegistry" binding:"omitempty,len=42"`
				PlatformFee     string `json:"platformFee" binding:"omitempty"`
				FeeRecipient    string `json:"feeRecipient" binding:"omitempty,len=42"`
			} `json:"marketplace"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("marketplace.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Marketplace, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid marketplace request in body", details)
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

		marketAddress := common.HexToAddress(cfg.CryptoConfig.BundleMarketplace)
		marketInstance, err := bundleMarketContract.NewFantomBundleMarketplace(marketAddress, client)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}

		if body.Marketplace.AddressRegistry != "" {
			address := common.HexToAddress(body.Marketplace.AddressRegistry)
			gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
			tx, err := marketInstance.UpdateAddressRegistry(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		if body.Marketplace.FeeRecipient != "" {
			address := common.HexToAddress(body.Marketplace.FeeRecipient)
			gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
			tx, err := marketInstance.UpdatePlatformFeeRecipient(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		if body.Marketplace.PlatformFee != "" {
			fee, err := strconv.ParseUint(body.Marketplace.PlatformFee, 10, 16)
			gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
			tx, err := marketInstance.UpdatePlatformFee(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, big.NewInt(int64(fee)))
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		// save marketplace
		marketplace := model.Marketplace{
			AddressRegistry: body.Marketplace.AddressRegistry,
			PlatformFee:     body.Marketplace.PlatformFee,
			FeeRecipient:    body.Marketplace.FeeRecipient,
		}
		err = h.marketplaceDB.SaveMarketplace(c.Request.Context(), &marketplace)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate collection name", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewMarketplaceResponse(&marketplace))
	})
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("v1/api")
	timeout := time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second
	v1.Use(middleware.RequestIDMiddleware(), middleware.TimeoutMiddleware(timeout))

	marketplaceV1 := v1.Group("marketplaces")
	// anonymous
	marketplaceV1.Use()
	{
		marketplaceV1.POST("/marketplace", h.saveMarketplace)
		marketplaceV1.POST("/bundlemarketplace", h.saveBundleMarketplace)
	}
}

func NewHandler(marketplaceDB marketplaceDB.MarketplaceDB) *Handler {
	return &Handler{
		marketplaceDB: marketplaceDB,
	}
}
