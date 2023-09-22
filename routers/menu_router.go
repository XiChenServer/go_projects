package routers

import "virus/api"

func (router RouterGroupe) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreateView)
	router.GET("menus", app.MenuListView)
	router.GET("menu_names", app.MenuNameList)
	router.PUT("menus/:id", app.MenuUpdateView)
	router.DELETE("menus", app.MenuRemoveView)
}
