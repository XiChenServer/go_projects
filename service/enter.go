package service

import (
	"virus/service/image_ser"
	"virus/service/user_ser"
)

type ServiceGroup struct {
	ImagerService image_ser.ImageService
	UserService   user_ser.UserService
}

var SrviceApp = new(ServiceGroup)
