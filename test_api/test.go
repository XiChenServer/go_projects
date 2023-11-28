package test_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestApi struct {
}

// @Tags 测试方法
// @Summary get请求测试方法
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /test/get_test [get]
func (TestApi) GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "你好",
	})
}

// @Tags 测试方法
// @Summary post请求测试方法
// @Param test formData string false "test"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /test/post_test [post]
func (TestApi) PostTest(c *gin.Context) {

	test := c.PostForm("test")
	if test == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "您发送的文件为空",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "我已经成功收到" + test,
	})
	fmt.Println(test)

}
