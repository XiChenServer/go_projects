package api

import "virus/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
