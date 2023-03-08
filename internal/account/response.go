package account

import (
	"rena-server-v2/internal/account/model"
	"time"
)

type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users      []User `json:"users"`
	UsersCount int64  `json:"usersCount"`
}

type User struct {
	Address   string    `json:"address"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserResponse(acc *model.Account) *UserResponse {
	return &UserResponse{
		User: User{
			Address:   acc.Address,
			Username:  acc.Username,
			Email:     acc.Email,
			Bio:       acc.Bio,
			Image:     acc.Image,
			CreatedAt: acc.CreatedAt,
		},
	}
}

// NewUsersResponse converts item models and total count to UsersResponse
func NewUsersResponse(accounts []*model.Account, total int64) *UsersResponse {
	var a []User
	for _, account := range accounts {
		a = append(a, NewUserResponse(account).User)
	}

	return &UsersResponse{
		Users:      a,
		UsersCount: total,
	}
}
