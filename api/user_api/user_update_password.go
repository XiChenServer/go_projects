package user_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models"
	"virus/models/res"
	"virus/utils/jwts"
	"virus/utils/pwd"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd"  binding:"required" msg:"请输入密码"`
	Pwd    string `json:"pwd"  binding:"required" msg:"请再次输入密码"`
}

func (UserApi) UserUpdatePassword(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if pwd.CheckPwd(user.Password, cr.OldPwd) {
		res.FailWithMessage("密码错误", c)
	}
	hashPwd := pwd.HashPwd(cr.Pwd)
	err = global.DB.Model(&user).Update("password", hashPwd).Error
	if err != nil {
		res.FailWithMessage("密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
	return
}
