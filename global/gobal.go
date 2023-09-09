package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"virus/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
