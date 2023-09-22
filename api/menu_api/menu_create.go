package menu_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models"
	"virus/models/ctype"
	"virus/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title"  binding:"required" msg:"请完善菜单名称" struct:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" struct:"path"`
	Slogan        string      `json:"slogan" struct:"slogan"`
	Abstract      ctype.Array `json:"abstract" struct:"abstract"`
	AbstractTime  int         `json:"abstract_time" struct:"abstract_time"`
	BannerTime    int         `json:"banner_time" struct:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请完善菜单序号" struct:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" struct:"-"`
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected //Model(models.MenuModel{}).Where("title = ? or path = ?", cr.Title, cr.Path).Find()
	if count > 0 {
		res.FailWithMessage("重复的菜单", c)
		return
	}
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menuModel).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}

	var menuBannerList []models.MenuBannerModel
	for _, sort := range cr.ImageSortList {
		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: sort.ImageID,
			Sort:     sort.Sort,
		})
	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功", c)
}
