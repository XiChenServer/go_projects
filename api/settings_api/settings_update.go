package settings_api

import (
	"github.com/gin-gonic/gin"
	"virus/config"
	"virus/core"
	"virus/global"
	"virus/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingUri
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if err == nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if err == nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if err == nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QQ = info
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBindJSON(&info)
		if err == nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.QiNiu = info
	case "jwt":
		var info config.Jwy
		err = c.ShouldBindJSON(&info)
		if err == nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.Jwy = info
	default:
		res.FailWithMessage("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
	res.OkWith(c)
}
