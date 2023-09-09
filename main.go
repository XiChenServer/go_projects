package main

import (
	"fmt"
	"virus/core"
	"virus/global"
)

func main() {
	core.InitConf()

	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
