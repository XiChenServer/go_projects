package service

import "virus/service/image_ser"

type ServiceGroup struct {
	ImagerService image_ser.ImageService
}

var SrviceApp = new(ServiceGroup)
