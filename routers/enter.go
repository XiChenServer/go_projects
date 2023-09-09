package routers

import (
	"github.com/gin-gonic/gin"
	"virus/global"
)

type RouterGroupe struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroupe{
		apiRouterGroup,
	}
	//系统配置api
	routerGroupApp.SettingsRouter()
	return router
}
