package images_api

import (
	"github.com/gin-gonic/gin"
	"virus/models"
	"virus/models/res"
	"virus/service/common"
	_ "virus/service/common"
)

func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//var p []models.BannerModel
	//count := global.DB.Find(&p).RowsAffected
	//fmt.Println(count)

	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	res.OkWithList(list, count, c)

	return
}
