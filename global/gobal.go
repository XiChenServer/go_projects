package global

import (
	"gorm.io/gorm"
	"virus/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
