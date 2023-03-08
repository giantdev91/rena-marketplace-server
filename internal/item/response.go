package item

import (
	collectionModel "rena-server-v2/internal/collection/model"
	itemModel "rena-server-v2/internal/item/model"
)

type TokenResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tokens []Item `json:"tokens"`
		Total  int64  `json:"total"`
	}
}

type ItemResponse struct {
	Item Item `json:"item"`
}

type ItemsResponse struct {
	Items      []Item `json:"items"`
	ItemsCount int64  `json:"itemsCount"`
}

type Item struct {
	ID                        uint    `json:"id"`
	ContentType               string  `json:"contentType"`
	ContractAddress           string  `json:"contractAddress"`
	ImageURL                  string  `json:"imageURL"`
	IsAppropriate             string  `json:"isAppropriate"`
	LastSalePrice             float32 `json:"lastSalePrice"`
	LastSalePriceInUSD        float32 `json:"lastSalePriceInUSD"`
	LastSalePricePaymentToken string  `json:"lastSalePricePaymentToken"`
	Liked                     int     `json:"liked"`
	Name                      string  `json:"name"`
	PaymentToken              string  `json:"paymentToken"`
	Price                     uint64  `json:"price"`
	PriceInUSD                float32 `json:"priceInUSD"`
	Supply                    int     `json:"supply"`
	ThumbnailPath             string  `json:"thumbnailPath"`
	TokenID                   uint    `json:"tokenId"`
	TokenType                 int     `json:"tokenType"`
	TokenURI                  string  `json:"tokenURI"`
	Creator                   string  `json:"creator"`
	Owner                     string  `json:"owner"`
	Collection                collectionModel.Collection
}

// NewItemsResponse converts item models and total count to ItemsResponse
func NewItemsResponse(items []*itemModel.Item, total int64) *ItemsResponse {
	var a []Item
	for _, item := range items {
		a = append(a, NewItemResponse(item).Item)
	}

	return &ItemsResponse{
		Items:      a,
		ItemsCount: total,
	}
}

// NewItemResponse converts item model to ItemResponse
func NewItemResponse(a *itemModel.Item) *ItemResponse {
	return &ItemResponse{
		Item: Item{
			ID:                        a.ID,
			ContractAddress:           a.ContractAddress,
			ContentType:               a.ContentType,
			ImageURL:                  a.ImageURL,
			IsAppropriate:             a.IsAppropriate,
			LastSalePrice:             a.LastSalePrice,
			LastSalePriceInUSD:        a.LastSalePriceInUSD,
			LastSalePricePaymentToken: a.LastSalePricePaymentToken,
			Liked:                     a.Liked,
			Name:                      a.Name,
			PaymentToken:              a.PaymentToken,
			Price:                     a.Price,
			PriceInUSD:                a.PriceInUSD,
			Supply:                    a.Supply,
			ThumbnailPath:             a.ThumbnailPath,
			TokenID:                   a.TokenID,
			TokenType:                 a.TokenType,
			TokenURI:                  a.TokenURI,
			Creator:                   a.Creator,
			Owner:                     a.Owner,
			Collection:                a.Collection,
		},
	}
}
