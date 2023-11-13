package router

import (
	"github.com/gin-gonic/gin"
	"go_ranking/controllers"
	"go_ranking/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default() //生成实例
	//创建请求
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	//r.GET("/url", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "Hello word")
	//})
	user := r.Group("/user")
	{

		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)
		user.POST("/add", controllers.UserController{}.AddUser)
		user.POST("/update", controllers.UserController{}.UpdateUser)
		user.POST("/list", controllers.UserController{}.GetList)
		user.DELETE("/delete", controllers.UserController{}.DeleteUser)
		user.POST("/list/test", controllers.UserController{}.GetUserListTest)
	}
	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}
	return r
}
