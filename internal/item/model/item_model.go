package model

import (
	collectionModel "rena-server-v2/internal/collection/model"
	"time"
)

type Item struct {
	ID                        uint      `gorm:"column:id"`
	ContentType               string    `gorm:"column:content_type"`
	ContractAddress           string    `gorm:"column:contract_address"`
	ImageURL                  string    `gorm:"column:image_url"`
	IsAppropriate             string    `gorm:"column:is_appropriate"`
	LastSalePrice             float32   `gorm:"column:last_sale_price"`
	LastSalePriceInUSD        float32   `gorm:"column:last_sale_price_in_usd"`
	LastSalePricePaymentToken string    `gorm:"column:last_sale_price_payment_token"`
	Liked                     int       `gorm:"column:liked"`
	Name                      string    `gorm:"column:name"`
	PaymentToken              string    `gorm:"column:payment_token"`
	Price                     uint64    `gorm:"column:price"`
	PriceInUSD                float32   `gorm:"column:price_in_usd"`
	Supply                    int       `gorm:"column:supply"`
	ThumbnailPath             string    `gorm:"column:thumbnail_path"`
	TokenID                   uint      `gorm:"column:token_id"`
	TokenType                 int       `gorm:"column:token_type"`
	TokenURI                  string    `gorm:"column:token_uri"`
	CollectionID              uint      `gorm:"column:collection_id"`
	Creator                   string    `gorm:"column:creator"`
	Owner                     string    `gorm:"column:owner"`
	CreatedAt                 time.Time `gorm:"column:created_at"`
	UpdatedAt                 time.Time `gorm:"column:updated_at"`
	DeletedAtUnix             int64     `gorm:"column:deleted_at_unix"`
	Collection                collectionModel.Collection
}
