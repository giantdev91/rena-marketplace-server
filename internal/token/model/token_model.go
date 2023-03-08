package model

import (
	"time"
)

type Token struct {
	ID            uint      `gorm:"column:id"`
	Token         string    `gorm:"column:token"`
	CreatedAt     time.Time `gorm:"column:created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at"`
	DeletedAtUnix int64     `gorm:"column:deleted_at_unix"`
}
