package router

import (
	"github.com/gin-gonic/gin"
	"go_ranking/controllers"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default() //生成实例
	//创建请求
	r.GET("/url", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello word")
	})
	user := r.Group("/user")
	{

		user.GET("/info", controllers.UserController{}.GetUserInfo)
		user.PUT("/put", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "Hello put")
		})
		user.POST("/post", controllers.UserController{}.GetList)
		user.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "/user/delete")
		})
	}
	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return r
}
