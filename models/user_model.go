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
	UserName string `gorm:"size:36" json:"user_name" `
	Password string `gorm:"size:128" json:"password"`
	AvatarID uint   `gorm:"size:256" json:"avatar_id"`
	//Avatar      ImageModel                    `json:"-"`
	Email          string           `gorm:"size:128" json:"email"`
	Tel            string           `json:"tel" gorm:"size:18"`
	Addr           string           `gorm:"size:64" json:"addr"`
	Token          string           `gorm:"size:64" json:"token"`
	IP             string           `gorm:"size:20" json:"IP"`
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`
	SignStatus     ctype.SignStatus ` gorm:"type=smallint(6)" json:"sign_status"`
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_model;joinForeignKey;UserID;JoinReferences:ArticleID" json:"-"`
}
