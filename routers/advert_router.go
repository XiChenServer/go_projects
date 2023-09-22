package routers

import "virus/api"

func (router RouterGroupe) AdvertRouter() {
	app := api.ApiGroupApp.AdverstApi
	router.POST("adverts", app.AdvertCreateView)
	router.GET("adverts", app.AdvertListView)
	router.PUT("adverts/:id", app.AdvertUpdateView)
	router.DELETE("adverts", app.AdvertRemoveView)
}
