package routers

import (
	"virus/api"
)

func (router RouterGroupe) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
	//router.GET("settings_email", settingsApi.SettingsEmailInfoView)
	//router.PUT("settings_email", settingsApi.SettingsEmailInfoUpdateView)
}
