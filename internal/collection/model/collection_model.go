package model

import (
	"time"
)

type Collection struct {
	ID              uint      `gorm:"column:id"`
	ContractAddress string    `gorm:"column:contract_address"`
	Name            string    `gorm:"column:name"`
	Slug            string    `gorm:"column:slug"`
	Description     string    `gorm:"column:description"`
	ImageURL        string    `gorm:"column:image_url"`
	Royalty         int       `gorm:"column:royalty"`
	FeeRecipient    string    `gorm:"column:fee_recipient"`
	Category        string    `gorm:"column:category"`
	Website         string    `gorm:"column:website"`
	Discord         string    `gorm:"column:discord"`
	TwitterUrl      string    `gorm:"column:twitter_url"`
	InstagramUrl    string    `gorm:"column:instagram_url"`
	MediumUrl       string    `gorm:"column:medium_url"`
	TelegramUrl     string    `gorm:"column:telegram_url"`
	ContactEmail    string    `gorm:"column:contact_email"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}
