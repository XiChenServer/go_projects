package flag

import (
	"fmt"
	"virus/global"
	"virus/models"
	"virus/models/ctype"
	"virus/utils/pwd"
)

func CreateUser(permissions string) {
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scan(&password)
	fmt.Printf("请再次输入密码：")
	fmt.Scan(&rePassword)
	fmt.Printf("请输入邮箱：")
	fmt.Scanln(&email)
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	if password != rePassword {
		global.Log.Error("两次密码不一致，请重新输入")
		return
	}
	hashPwd := pwd.HashPwd(password)

	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	avatar := "/uploads/avatar/default.png"

	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Info("用户%s创建成功", userName)
}
