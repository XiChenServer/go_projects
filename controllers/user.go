package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go_ranking/models"
	"strconv"
)

type UserController struct {
}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "两次输入密码不同")
		return
	}
	user, _ := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "该用户已经存在")
		return
	}
	id, _ := models.AddUser(username, EncryMd5(password))
	ReturnSuccess(c, 0, "注册成功", id, 1)
}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()
	data := UserApi{Id: user.Id, Username: user.Username}
	ReturnSuccess(c, 0, "登录成功", data, 1)
}
