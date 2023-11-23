package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "请输入当前页，默认第一页"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	size, err := strconv.Atoi(c.DefaultPostForm("size", define.DefaultSize))
	if err != nil {
		log.Println("getproblemlist size parse error:", err)
		return
	}
	page, err := strconv.Atoi(c.DefaultPostForm("page", define.DefaultPage))
	if err != nil {
		log.Println("getproblemlist page parse error:", err)
		return
	}
	//page == 1 == > offset 0
	page = (page - 1) * size
	var count int64
	keywork := c.Query("keyword")
	list := make([]*models.Problem, 0)
	tx := models.GetProblemList(keywork)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("get problem list error:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"data":  &list,
			"count": count,
		},
	})
}
