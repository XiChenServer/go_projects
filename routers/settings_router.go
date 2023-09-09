package routers

import (
	"virus/api"
)

func (router RouterGroupe) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingInfoView)

}
