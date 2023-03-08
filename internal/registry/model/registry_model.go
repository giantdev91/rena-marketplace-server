package model

import (
	"time"
)

type Registry struct {
	ID              uint      `gorm:"column:id"`
	ContractType    string    `gorm:"column:contract_type"`
	ContractAddress string    `gorm:"column:contract_address"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	DeletedAtUnix   int64     `gorm:"column:deleted_at_unix"`
}
