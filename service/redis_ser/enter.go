package redis_ser

import (
	"github.com/gin-gonic/gin"
	"time"
	"virus/global"
	"virus/utils"
)

const prefix = "logout_"

func Logout(c *gin.Context, token string, diff time.Duration) error {
	err := global.Redis.Set(c, prefix+token, "", diff).Err()
	return err
}
func CheckLogout(c *gin.Context, token string) bool {
	keys := global.Redis.Keys(c, prefix+"*").Val()
	if utils.InList(prefix+token, keys) {
		return true
	}
	return false

}
