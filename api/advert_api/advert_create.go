package advert_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models"
	"virus/models/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" struct:"title"`
	Href   string `json:"href"  binding:"required,url" msg:"非法跳转连接" struct:"href"`
	Images string `json:"images" binding:"required,url" msg:"非法图片地址" struct:"images"`
	IsShow bool   `json:"is_show"  msg:"请选择是否展示" struct:"is_show"`
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest     true "表示多个参数"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("广告已存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)
}
