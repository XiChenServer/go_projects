package models

type MenuImageModel struct {
	MenuID      uint        `json:"menu_id"`
	MenuModel   MenuModel   `gorm:"foreignKey:MenuId"`
	BanneriD    uint        `json:"banner_id"`
	BannerModel BannerModel `gorm:"foreignKey:BanneriD"`
	Sort        int         `gorm:"size:18" json:"sort"`
}
