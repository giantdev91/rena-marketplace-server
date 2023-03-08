package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Account struct {
	ID        uint      `gorm:"column:id"`
	Address   string    `gorm:"column:address"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Bio       string    `gorm:"column:bio"`
	Image     string    `gorm:"column:image"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (a Account) String() string {
	return fmt.Sprintf("Account{id:%d, address:%s, username:%s, bio:%s, image:%s, createdAt:%v, updatedAt:%v",
		a.ID, a.Address, a.Username, a.Bio, a.Image, a.CreatedAt, a.UpdatedAt)
}

func (a *Account) UnmarshalJSON(b []byte) error {
	var tmp struct {
		ID       uint   `json:"id"`
		Address  string `json:"address"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}
	a.ID = tmp.ID
	a.Address = tmp.Address
	a.Username = tmp.Username
	a.Email = tmp.Email
	return nil
}
