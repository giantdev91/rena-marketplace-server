package item

import (
	"fmt"
	"net/http"
	collectionDB "rena-server-v2/internal/collection/database"
	"rena-server-v2/internal/config"
	"rena-server-v2/internal/database"
	itemDB "rena-server-v2/internal/item/database"
	"rena-server-v2/internal/item/model"
	"rena-server-v2/internal/middleware"
	"rena-server-v2/internal/middleware/handler"
	"rena-server-v2/pkg/logging"
	"rena-server-v2/pkg/validate"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	collectionDB collectionDB.CollectionDB
	itemDB       itemDB.ItemDB
}

// saveItem handles POST /v1/api/items
func (h *Handler) saveItem(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Item struct {
				ContractAddress string `json:"contractAddress" binding:"required,min=40"`
				TokenID         uint   `json:"tokenID" binding:"required"`
				TokenURI        string `json:"tokenURI" binding:"required"`
				ImageURL        string `json:"imageURL" binding:"omitempty"`
				Name            string `json:"name" binding:"required"`
				Description     string `json:"description" binding:"required"`
				Creator         string `json:"creator" binding:"required"`
			} `json:"item"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("item.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Item, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid collection request in body", details)
		}

		// cfg, err := config.Load("")
		// if err != nil {
		// 	return handler.NewInternalErrorResponse(err)
		// }

		// client, err := ethclient.Dial(cfg.CryptoConfig.Network)
		// if err != nil {
		// 	return handler.NewInternalErrorResponse(err)
		// }

		// privateKey, err := crypto.HexToECDSA(cfg.CryptoConfig.PrivateKey)
		// if err != nil {
		// 	return handler.NewInternalErrorResponse(err)
		// }

		// publicKey := privateKey.Public()
		// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		// if !ok {
		// 	return handler.NewInternalErrorResponse(err)
		// }

		// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		// if err != nil {
		// 	return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		// }

		// chainID, err := client.ChainID(context.Background())
		// if err != nil {
		// 	return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		// }

		// auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		// if err != nil {
		// 	return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		// }
		// auth.Nonce = big.NewInt(int64(nonce))
		// auth.Value = big.NewInt(0) // in wei

		// marketAddress := common.HexToAddress(cfg.CryptoConfig.Marketplace)
		// marketInstance, err := marketContract.NewFantomMarketplace(marketAddress, client)
		// if err != nil {
		// 	return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		// }

		// address := common.HexToAddress(body.Item.ContractAddress)
		// tokenID := big.NewInt(int64(body.Item.TokenID))
		// gasPrice, _ := client.SuggestGasPrice(c.Request.Context())
		// tx, err := marketInstance.ListItem(&bind.TransactOpts{
		// 	From:     auth.From,
		// 	Signer:   auth.Signer,
		// 	GasPrice: gasPrice,
		// }, address, tokenID, big.NewInt(1), common.HexToAddress(""), big.NewInt(0), big.NewInt(0))
		// if err != nil {
		// 	return handler.NewErrorResponse(http.StatusInternalServerError, handler.InternalServerError, err.Error(), nil)
		// }
		// fmt.Println(tx)

		collection, _ := h.collectionDB.FindCollectionByAddress(c.Request.Context(), body.Item.ContractAddress)

		// save item
		item := model.Item{
			ContentType:               "",
			ContractAddress:           body.Item.ContractAddress,
			ImageURL:                  body.Item.ImageURL,
			IsAppropriate:             "",
			LastSalePrice:             0.0,
			LastSalePriceInUSD:        0.0,
			LastSalePricePaymentToken: "",
			Liked:                     0,
			Name:                      body.Item.Name,
			PaymentToken:              "",
			Price:                     0.0,
			PriceInUSD:                0.0,
			Supply:                    1,
			ThumbnailPath:             "",
			TokenID:                   body.Item.TokenID,
			TokenType:                 721,
			TokenURI:                  body.Item.TokenURI,
			CollectionID:              collection.ID,
			Creator:                   body.Item.Creator,
			Owner:                     body.Item.Creator,
		}
		err := h.itemDB.SaveItem(c.Request.Context(), &item)
		if err != nil {
			if database.IsKeyConflictErr(err) {
				return handler.NewErrorResponse(http.StatusConflict, handler.DuplicateEntry, "duplicate collection name", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		item.Collection = *collection
		return handler.NewSuccessResponse(http.StatusCreated, NewItemResponse(&item))
	})
}

// updateItem handles POST /v1/api/items/update
func (h *Handler) updateItem(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type RequestBody struct {
			Item struct {
				ContractAddress string `json:"contractAddress" binding:"required,len=42"`
				TokenID         uint   `json:"tokenId" binding:"required"`
				Quantity        int    `json:"quantity" binding:"required"`
				PaymentToken    string `json:"paymentToken" binding:"omitempty"`
				PricePerItem    string `json:"pricePerItem" binding:"required"`
				StartingTime    string `json:"startingTime" binding:"required"`
			} `json:"item"`
		}
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Errorw("item.handler.register failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&body.Item, "json", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid item request in body", details)
		}

		item, err := h.itemDB.FindItemByAddress(c.Request.Context(), body.Item.ContractAddress, body.Item.TokenID)
		if err != nil {
			if database.IsRecordNotFoundErr(err) {
				return handler.NewErrorResponse(http.StatusNotFound, handler.NotFoundEntity, "not found item", nil)
			}
			return handler.NewInternalErrorResponse(err)
		}
		pricePerItem, err := strconv.ParseUint(body.Item.PricePerItem, 10, 64)
		if err != nil {
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidBodyValue, "invalid item request in body", err)
		}
		item.PaymentToken = body.Item.PaymentToken
		item.Price = pricePerItem
		item.Supply = body.Item.Quantity

		// update item
		err = h.itemDB.UpdateItem(c.Request.Context(), item)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusCreated, NewItemResponse(item))
	})
}

// items handles GET /v1/api/items
func (h *Handler) items(c *gin.Context) {
	handler.HandleRequest(c, func(c *gin.Context) *handler.Response {
		logger := logging.FromContext(c)

		type QueryParameter struct {
			Collection []uint `form:"collection" binding:"omitempty,dive,max=10"`
			Creator    string `form:"creator" binding:"omitempty,len=42"`
			Owner      string `form:"owner" binding:"omitempty,len=42"`
			Payment    string `form:"payment" binding:"omitempty"`
			PriceMin   string `form:"price_min,default=0" binding:"omitempty"`
			PriceMax   string `form:"price_max,default=0" binding:"omitempty"`
			Offset     uint   `form:"offset,default=0" binding:"numeric"`
			Limit      uint   `form:"limit,default=6" binding:"numeric"`
		}
		var query QueryParameter
		if err := c.ShouldBindQuery(&query); err != nil {
			logger.Errorw("item.handler.items failed to bind", "err", err)
			var details []*validate.ValidationErrDetail
			if vErrs, ok := err.(validator.ValidationErrors); ok {
				details = validate.ValidationErrorDetails(&query, "form", vErrs)
			}
			return handler.NewErrorResponse(http.StatusBadRequest, handler.InvalidUriValue, "invalid item request in query", details)
		}

		fmt.Println(query)
		priceMin, err := strconv.ParseUint(query.PriceMin, 10, 64)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		priceMax, err := strconv.ParseUint(query.PriceMax, 10, 64)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		criteria := itemDB.IterateItemCriteria{
			Collections: query.Collection,
			Creator:     query.Creator,
			Owner:       query.Owner,
			PriceMin:    priceMin,
			PriceMax:    priceMax,
			Offset:      query.Offset,
			Limit:       query.Limit,
		}
		items, total, err := h.itemDB.FindItems(c.Request.Context(), criteria)
		if err != nil {
			return handler.NewInternalErrorResponse(err)
		}
		return handler.NewSuccessResponse(http.StatusOK, NewItemsResponse(items, total))
	})
}

func RouteV1(cfg *config.Config, h *Handler, r *gin.Engine, auth *jwt.GinJWTMiddleware) {
	v1 := r.Group("v1/api")
	timeout := time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second
	v1.Use(middleware.RequestIDMiddleware(), middleware.TimeoutMiddleware(timeout))

	itemV1 := v1.Group("items")
	// anonymous
	itemV1.Use()
	{
		itemV1.GET("", h.items)
		itemV1.POST("", h.saveItem)
		itemV1.PUT("", h.updateItem)
	}
}

func NewHandler(itemDB itemDB.ItemDB, collectionDB collectionDB.CollectionDB) *Handler {
	return &Handler{
		itemDB:       itemDB,
		collectionDB: collectionDB,
	}
}
