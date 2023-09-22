package main

import (
	"virus/core"
	_ "virus/docs"
	"virus/flag"
	"virus/global"
	"virus/routers"
)

// @title           gvb_server API文档
// @version         1.0
// @description     gvb_server API文档
// @host      127.0.0.1:8080
// @BasePath  /
func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	router := routers.InitRouter()

	err := router.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
