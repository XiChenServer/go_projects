package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	ReturnSuccess(c, 0, "success", "user info", 1)
}

func (u UserController) GetList(c *gin.Context) {
	ReturnError(c, 4004, "list没有相关信息")
}
