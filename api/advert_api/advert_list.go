package advert_api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"virus/models"
	"virus/models/res"
	"virus/service/common"
)

// AdvertListView 广告列表
// @Tags 广告管理
// @Summary 广告列表
// @Description 广告列表
// @Param data query models.PageInfo     false "查询参数"
// @Router /api/adverts [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
	})
	res.OkWithList(list, count, c)
}
