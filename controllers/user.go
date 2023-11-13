package controllers

import (
	"github.com/gin-gonic/gin"
	"go_ranking/models"
	"strconv"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")
	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 0, name, user, 1)
}

func (u UserController) GetList(c *gin.Context) {

	//logger.Write("日志信息", "user")
	//ReturnError(c, 4004, "list没有相关信息")
}
func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("name", "")

	id, err := models.AddUser(username)
	if err != nil {
		ReturnError(c, 4002, "保存错误")
		return
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)
}

func (u UserController) UpdateUser(c *gin.Context) {
	username := c.DefaultPostForm("name", "")
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.UpdateUser(id, username)
	if err != nil {
		ReturnError(c, 4002, "修改错误")
		return
	}
	ReturnSuccess(c, 0, "修改成功", id, 1)
}
func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(id)
	if err != nil {
		ReturnError(c, 4002, "删除错误")
		return
	}
	ReturnSuccess(c, 0, "删除成功", id, 1)
}
func (u UserController) GetUserListTest(c *gin.Context) {
	idStr := c.DefaultPostForm("id", "")
	id, _ := strconv.Atoi(idStr)
	users, err := models.GetUserListTest(id)
	if err != nil {
		ReturnError(c, 4002, "获取错误")
		return
	}
	ReturnSuccess(c, 0, "获取成功", users, 1)
}
