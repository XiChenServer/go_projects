package models

import (
	"fmt"
	"go_ranking/dao"
)

type User struct {
	Id       int
	Username string
}

func (User) TableName() string {
	return "user"
}
func GetUserTest(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username string) (int, error) {
	user := User{Username: username}
	fmt.Println(username)
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

func UpdateUser(id int, username string) error {
	err := dao.Db.Model(&User{}).Where("id = ?", id).Update("username", username).Error
	return err
}
func DeleteUser(id int) error {
	err := dao.Db.Delete(&User{}, id).Error
	return err
}
func GetUserListTest(id int) ([]User, error) {
	var users []User
	err := dao.Db.Where("id < ?", id).Find(&users).Error
	return users, err

}
