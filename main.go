package main

import (
	"gin_chat/models"
	"gin_chat/router"
	"gin_chat/utils"
	"github.com/spf13/viper"
	"time"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	utils.DB.AutoMigrate(&models.UserBasic{}, &models.Community{}, &models.Contact{}, &models.GroupBasic{}, &models.Message{})
	r := router.Router()
	r.Run(":8081")
}

// 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
