package main

import (
	"gin_chat/models"
	"gin_chat/router"
	"gin_chat/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	utils.DB.AutoMigrate(&models.UserBasic{}, &models.Community{}, &models.Contact{}, &models.GroupBasic{}, &models.Message{})
	r := router.Router()
	r.Run(":8081")
}
