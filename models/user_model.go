package models

import "virus/models/ctype"

type Role int

const (
	PermissionAdmin       Role = 1
	PermissionUser        Role = 2
	PermissionVisitor     Role = 3
	PermissionDisableUser Role = 4
)

type UserModel struct {
	MODEL
	NickName string `gorm:"size:36" json:"nick_name"`
	UserName string `json:"user_name" gorm:"size:36"`
	Password string `gorm:"size:128" json:"password"`
	AvatarID uint   `gorm:"size:256" json:"avatar_id"`
	//Avatar      ImageModel                    `json:"-"`
	Email          string           `gorm:"size:128" json:"email"`
	Tel            string           `json:"tel" gorm:"size:18"`
	Addr           string           `gorm:"size:64" json:"addr"`
	Token          string           `gorm:"size:64" json:"token"`
	IP             string           `gorm:"size:20" json:"IP"`
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus     ctype.SignStatus `json:"sign_status" gorm:"type=smallint(6)"`
	ArticleMode    []ArticleModel   `gorm:"foreignKey" json:"-"`
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;joinForeignKey;AuthID;JoinReferences:ArticleID" json:"-"`
}
