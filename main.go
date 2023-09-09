package main

import (
	"virus/core"
	"virus/global"
	"virus/routers"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	router := routers.InitRouter()
	router.Run(global.Config.System.Addr())
}
