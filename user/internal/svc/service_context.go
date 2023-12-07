package svc

import (
	"gorm.io/gorm"
	"iot-platform/models"
	"iot-platform/user/internal/config"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		DB:     models.DB,
	}
}
