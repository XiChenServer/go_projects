package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type VoteController struct {
}

func (v VoteController) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerStr, _ := strconv.Atoi(playerIdStr)
	if userId == 0 || playerStr == 0 {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}

}
