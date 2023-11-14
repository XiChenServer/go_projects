package main

import (
	"go_ranking/router"
)

func main() {

	r := router.Router()

	//defer 后执行
	//recover可以让崩溃的程序回复过来
	//panic让程序崩溃
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获异常:", err)
	//	}
	//}()
	//panic("11")
	r.Run(":9999")
}
