package user_api

import (
	"github.com/gin-gonic/gin"
	"virus/models"
	"virus/models/ctype"
	"virus/models/res"
	"virus/service/common"
	"virus/utils/desens"
	"virus/utils/jwts"
)

func (UserApi) UserListView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	//fmt.Println(token)
	if token == "" {
		res.FailWithMessage("没携带token", c)
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token错误", c)
		return
	}
	if ctype.Role(claims.Role) == ctype.PermissionAdmin {

	}
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)

}
