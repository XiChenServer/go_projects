package common

import (
	"gorm.io/gorm"
	"virus/global"
	"virus/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc"
	}
	//var p []models.BannerModel
	//i := global.DB.Find(&p).RowsAffected
	//fmt.Println(i)
	count = DB.Debug().Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	//fmt.Println(list)
	err = DB.Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
