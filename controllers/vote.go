package controllers

import (
	"github.com/gin-gonic/gin"
	"go_ranking/cache"
	"go_ranking/models"
	"strconv"
)

type VoteController struct {
}

func (v VoteController) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.Atoi(userIdStr)
	playerId, _ := strconv.Atoi(playerIdStr)
	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	user, _ := models.GetUserInfo(userId)
	if user.Id == 0 {
		ReturnError(c, 4001, "投票用户不存在")
		return
	}
	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "参赛选手不存在")
		return
	}
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "已经投票")
		return
	}
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		models.UpdatePlayerScore(playerId)
		var redisKey string
		redisKey = "ranking:" + strconv.Itoa(player.Aid)
		//获取redis信息
		cache.Rdb.ZIncrBy(cache.Rctx, redisKey, 1, strconv.Itoa(playerId))

		ReturnSuccess(c, 0, "投票成功", rs, 1)
		return
	}
	ReturnError(c, 4004, "请联系管理员")
	return
}
