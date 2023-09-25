package user_ser

import (
	"github.com/gin-gonic/gin"
	"time"
	"virus/service/redis_ser"
	"virus/utils/jwts"
)

type UserService struct {
}

func (UserService) Logout(c *gin.Context, claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(c, token, diff)

}
