package model

import (
	"time"
)

type Marketplace struct {
	ID              uint      `gorm:"column:id"`
	AddressRegistry string    `gorm:"column:address_registry"`
	PlatformFee     string    `gorm:"column:platform_fee"`
	FeeRecipient    string    `gorm:"column:fee_recipient"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	DeletedAtUnix   int64     `gorm:"column:deleted_at_unix"`
}
