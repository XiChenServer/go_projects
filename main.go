package main

import (
	"virus/core"
	"virus/flag"
	"virus/global"
	"virus/routers"
)

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
