package settings_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models/res"
)

func (SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OkWithData(global.Config.Email, c)
}
