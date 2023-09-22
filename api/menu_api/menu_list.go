package menu_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models"
	"virus/models/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banners"`
}

func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)

	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menus []MenuResponse
	for _, model := range menuList {
		var banners = make([]Banner, 0)
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.BannerID,
				Path: banner.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menus, c)
	return
}
