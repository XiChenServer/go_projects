package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go_ranking/config"
	"go_ranking/pkg/logger"
	"time"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.Mysqldb)
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error": err.Error()})
	}
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"database error": Db.Error})
	}
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
	//logger.Info(map[string]interface{}{"Database connection initialized": err.Error()})
}
