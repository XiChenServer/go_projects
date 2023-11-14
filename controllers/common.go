package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"go_ranking/dao"
	"go_ranking/models"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
type JsonErrstruct struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
}

func InitSqlTable() {
	dao.Db.AutoMigrate(&models.User{})

	dao.Db.AutoMigrate(&models.Player{})

	dao.Db.AutoMigrate(&models.Vote{})
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int64) {
	json := &JsonStruct{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	c.JSON(200, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonErrstruct{Code: code, Msg: msg}
	c.JSON(200, json)
}
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
