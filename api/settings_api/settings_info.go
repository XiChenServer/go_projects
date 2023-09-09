package settings_api

import (
	"github.com/gin-gonic/gin"
	"virus/models/res"
)

func (SettingsApi) SettingInfoView(c *gin.Context) {
	res.Ok(map[string]string{}, "sf", c)
	c.JSON(200, gin.H{"msg": "222"})
}
