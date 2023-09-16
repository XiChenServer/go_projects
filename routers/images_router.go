package routers

import "virus/api"

func (router RouterGroupe) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("images", app.ImageUploadView)
	router.GET("images", app.ImageListView)
	router.DELETE("images", app.ImageRemoveView)
	router.PUT("images", app.ImageUpdateView)
}
