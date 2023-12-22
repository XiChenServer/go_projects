package main

import (
	"gin_chat/models"
	"gin_chat/router"
	"gin_chat/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()
	utils.DB.AutoMigrate(&models.UserBasic{})
	r := router.Router()
	r.Run()
}
