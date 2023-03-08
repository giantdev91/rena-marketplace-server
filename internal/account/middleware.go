package account

import (
	"net/http"
	accountDB "rena-server-v2/internal/account/database"
	"rena-server-v2/internal/account/model"
	"rena-server-v2/internal/config"
	"rena-server-v2/pkg/logging"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

type signIn struct {
	User struct {
		Address string `form:"address" json:"address" binding:"required"`
	} `json:"user"`
}

func CurrentUser(c *gin.Context) (*model.Account, bool) {
	data, ok := c.Get(identityKey)
	if !ok {
		return nil, false
	}
	acc, ok := data.(*model.Account)
	return acc, ok
}

func MustCurrentUser(c *gin.Context) *model.Account {
	acc, ok := CurrentUser(c)
	if ok {
		return acc
	}
	panic("no account in gin.Context")
}

func NewAuthMiddleware(cfg *config.Config, accountDB accountDB.AccountDB) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(cfg.JwtConfig.Secret),
		Timeout:     time.Duration(cfg.JwtConfig.SessionTime) * time.Millisecond,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.Account); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			address := claims[identityKey].(string)
			logging.FromContext(c).Info("middleware.jwt.IdentityHandler", "address", address)
			acc, _ := accountDB.FindByAddress(c.Request.Context(), address)
			return acc
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var req signIn
			if err := c.ShouldBindJSON(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			logging.FromContext(c).Info("middleware.jwt.Authenticator", "address", req.User.Address)

			acc, err := accountDB.FindByAddress(c.Request.Context(), req.User.Address)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			return &model.Account{
				ID:       acc.ID,
				Address:  acc.Address,
				Username: acc.Username,
				Email:    acc.Email,
				Bio:      acc.Bio,
				Image:    acc.Image,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			logging.FromContext(c).Info("middleware.jwt.Authorizator", "data", data)
			if _, ok := data.(*model.Account); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			logging.FromContext(c).Info("middleware.jwt.Unauthorized", "code", code, "message", message)
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"code":   code,
				"token":  token,
				"expire": expire,
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}
