package api

import (
	"virus/api/advert_api"
	"virus/api/images_api"
	"virus/api/menu_api"
	"virus/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdverstApi  advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
