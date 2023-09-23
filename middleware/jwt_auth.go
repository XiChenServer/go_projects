package middleware

import (
	"github.com/gin-gonic/gin"
	"virus/models/res"
	"virus/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		//fmt.Println(token)
		if token == "" {
			res.FailWithMessage("没携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
