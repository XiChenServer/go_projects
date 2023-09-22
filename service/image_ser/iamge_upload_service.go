package image_ser

import (
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"strings"
	"virus/global"
	"virus/models"
	"virus/models/ctype"
	"virus/plugins/qiniu"
	"virus/utils"
)

var (
	WhiteImageList = []string{
		"jpg",
		"png",
		"jpeg",
		"ico",
		"tiff",
		"gif",
		"svg",
		"webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename

	baseList := global.Config.Upload.Path
	filePath := path.Join(baseList, file.Filename)
	res.FileName = filePath

	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		res.Msg = "文件非法"
		return
	}

	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("图片大小超过设定大小，设定大小为: %dMB，当前文件大小为： %fMB", global.Config.Upload.Size, size)
		return
	}
	//读取文件内容，hash
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	//查找是否含有
	var bannerModel models.BannerModel
	err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		res.Msg = "图片已存在"
		res.FileName = bannerModel.Path
		return
	}
	fileType := ctype.Local
	res.Msg = "图片上传成功"
	res.IsSuccess = true
	//fmt.Println(global.Config.QiNiu.Enable)
	if global.Config.QiNiu.Enable {
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error(err)
			res.Msg = err.Error()
			return
		}
		res.FileName = filePath
		res.Msg = "上传七牛成功"
		fileType = ctype.QiNiu
	}

	//图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})
	return res
}
