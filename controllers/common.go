package controllers

import "github.com/gin-gonic/gin"

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
