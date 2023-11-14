package controllers

import (
	"github.com/gin-gonic/gin"
	"go_ranking/models"
	"strconv"
)

type PlayerController struct {
}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid)
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}

	ReturnSuccess(c, 0, "success", rs, 1)
}
