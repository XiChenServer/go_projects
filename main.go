package main

import (
	"fmt"
	"virus/core"
	"virus/global"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()
	global.Log.Warnln("dfgg")
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
