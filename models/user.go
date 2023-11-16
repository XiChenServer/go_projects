package models

import (
	"go_ranking/dao"
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	AddTime    int64  `json:"addTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}
func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func GetUserInfo(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err

}
