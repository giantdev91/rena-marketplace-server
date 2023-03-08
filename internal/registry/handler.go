package registry

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
	registryDB "rena-server-v2/internal/registry/database"
	"rena-server-v2/internal/registry/model"
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

	registryContract "rena-server-v2/internal/contract/registry"
)

type Handler struct {
	registryDB registryDB.RegistryDB
}

// saveRegistry handles POST /v1/api/registry
func (h *Handler) saveRegistry(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Registry struct {
				ContractType    string `json:"contractType" binding:"required,oneof=rena auction marketplace bundlemarketplace factory privatefactory artfactory privateartfactory tokenregistry pricefeed"`
				ContractAddress string `json:"contractAddress" binding:"required,len=42"`
			} `json:"registry"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("registry.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Registry, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid collection request in body", details)
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

		registryAddress := common.HexToAddress(cfg.CryptoConfig.AddressRegistry)
		registryInstance, err := registryContract.NewFantomAddressRegistry(registryAddress, client)
		if err != nil {
			return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		}

		address := common.HexToAddress(body.Registry.ContractAddress)
		gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
		if body.Registry.ContractType == "rena" {
			tx, err := registryInstance.UpdateArtion(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "auction" {
			tx, err := registryInstance.UpdateAuction(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "marketplace" {
			tx, err := registryInstance.UpdateMarketplace(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "bundlemarketplace" {
			tx, err := registryInstance.UpdateBundleMarketplace(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "factory" {
			tx, err := registryInstance.UpdateNFTFactory(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "privatefactory" {
			tx, err := registryInstance.UpdateNFTFactoryPrivate(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "artfactory" {
			tx, err := registryInstance.UpdateArtFactory(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "privateartfactory" {
			tx, err := registryInstance.UpdateArtFactoryPrivate(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "tokenregistry" {
			tx, err := registryInstance.UpdateTokenRegistry(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		} else if body.Registry.ContractType == "pricefeed" {
			tx, err := registryInstance.UpdatePriceFeed(&bind.TransactOpts{
				From:     auth.From,
				Signer:   auth.Signer,
				GasPrice: gasPrice,
			}, address)
			if err != nil {
				return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
			}
			fmt.Println(tx)
		}

		// save registry
		registry := model.Registry{
			ContractType:    body.Registry.ContractType,
			ContractAddress: body.Registry.ContractAddress,
		}
		err = h.registryDB.SaveRegistry(c.Request.Context(), &registry)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate contract address", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewRegistryResponse(&registry))
	})
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("v1/api")
	timeout := time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second
	v1.Use(middleware.RequestIDMiddleware(), middleware.TimeoutMiddleware(timeout))

	registryV1 := v1.Group("registries")
	// anonymous
	registryV1.Use()
	{
		registryV1.POST("", h.saveRegistry)
	}
}

func NewHandler(registryDB registryDB.RegistryDB) *Handler {
	return &Handler{
		registryDB: registryDB,
	}
}
