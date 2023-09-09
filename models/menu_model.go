package models

import "virus/models/ctype"

type MenuModel struct {
	MODEL
	MenuTitle    string        `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string        `gorm:"size:32" json:"menu_title_en"`
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                     // slogan
	Abstract     ctype.Array   `gorm:"type:varchar(255)" json:"abstract"`                                                         //司介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;joinReferences:BannerID" json:"banners"` // 菜手的图片列来
	BannerTime   int           `json:"banner_time"`                                                                               //菜羊四片的切时间为日不切
	Sort         int           `gorm:"size:10" json:"sort"`                                                                       //菜单的暖序
}
