package routers

import (
	"virus/api"
	"virus/middleware"
)

func (router RouterGroupe) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
}
