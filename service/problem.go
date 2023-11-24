package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
// @Param category_identity query string false "category_identity"
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
	categoryIdentity := c.Query("category_identity")
	list := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keywork, categoryIdentity)
	err = tx.Omit("content").Count(&count).Offset(page).Limit(size).Find(&list).Error

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

// GetProblemList
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string false "problem identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "问题唯一标识不能为空",
		})
		return
	}
	data := new(models.ProblemBasic)
	err := models.DB.Where("identity = ?", identity).
		Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "当前问题不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get ProblemDetail Error:" + err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"data": data,
	})
}
