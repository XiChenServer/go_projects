package images_api

import (
	"github.com/gin-gonic/gin"
	"io/fs"
	"os"
	"virus/global"
	"virus/models/res"
	"virus/service"
	"virus/service/image_ser"
)

func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}

	//判断路径是否存在
	baseList := global.Config.Upload.Path
	_, err = os.ReadDir(baseList)
	if err != nil {
		err = os.MkdirAll(baseList, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	//不存在就创建
	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		serviceRes := service.SrviceApp.ImagerService.ImageUploadService(file)
		if !serviceRes.IsSuccess {
			resList = append(resList, serviceRes)
			continue
		}
		if !global.Config.QiNiu.Enable {
			err = c.SaveUploadedFile(file, serviceRes.FileName)
			if err != nil {
				global.Log.Error(err)
				serviceRes.Msg = err.Error()
				serviceRes.IsSuccess = false
				resList = append(resList, serviceRes)
				continue
			}
		}
		resList = append(resList, serviceRes)
	}
	res.OkWithData(resList, c)
}
