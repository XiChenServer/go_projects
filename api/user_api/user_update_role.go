package user_api

import (
	"github.com/gin-gonic/gin"
	"virus/global"
	"virus/models"
	"virus/models/ctype"
	"virus/models/res"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nick_name"`
	UserID   uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMessage("用户id错误，用户不存在", c)
		return
	}
	err = global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		res.FailWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}
