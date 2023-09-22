package models

import "virus/models/ctype"

// 菜单的路径可以是/path,也可以是路由别名
type MenuModel struct {
	MODEL
	Title        string        `gorm:"size:32" json:"title"`
	Path         string        `gorm:"size:32" json:"path"`
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                     // slogan
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                               //司介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` // 菜手的图片列来
	BannerTime   int           `json:"banner_time"`                                                                               //菜羊四片的切时间为日不切
	Sort         int           `gorm:"size:10" json:"sort"`
	//菜单的暖序
}
