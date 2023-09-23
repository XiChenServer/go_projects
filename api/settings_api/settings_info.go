package settings_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models/res"
)

type SettingUri struct {
	Name string `uri:"name"`
}

//var SettingsMap map[string]

func (SettingsApi) SettingInfoView(c *gin.Context) {

	var cr SettingUri
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwy":
		res.OkWithData(global.Config.Jwy, c)
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}

}
