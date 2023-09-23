package routers

import "virus/api"

func (router RouterGroupe) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", app.UserListView)
}
