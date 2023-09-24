package routers

import (
	"virus/api"
	"virus/middleware"
)

func (router RouterGroupe) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("email_login", app.EmailLoginView)
	router.GET("users", middleware.JwtAuth(), app.UserListView)
	router.PUT("user_role", middleware.JwtAuth(), app.UserUpdateRoleView)
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePassword)
}
