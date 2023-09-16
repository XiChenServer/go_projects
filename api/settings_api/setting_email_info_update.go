package settings_api

import (
	"github.com/gin-gonic/gin"
	"virus/config"
	"virus/core"
	"virus/global"
	"virus/models/res"
)

func (SettingsApi) SettingsEmailInfoUpdateView(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.Email = cr

	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWith(c)
}
