package controllers

import (
	"github.com/gin-gonic/gin"
	"go_ranking/cache"
	"go_ranking/models"
	"strconv"
	"time"
)

type PlayerController struct {
}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
	rs, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		ReturnError(c, 4004, "没有相关信息")
		return
	}

	ReturnSuccess(c, 0, "success", rs, 1)
}
func (p PlayerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	//定义rediskey
	var redisKey string
	redisKey = "ranking:" + aidStr
	//获取redis信息
	rs, err := cache.Rdb.ZRevRange(cache.Rctx, redisKey, 0, -1).Result()
	//if err == redis.Nil {
	if err == nil && len(rs) > 0 {
		var players []models.Player
		for _, value := range rs {
			id, _ := strconv.Atoi(value)
			rsInfo, _ := models.GetPlayerInfo(id)
			if rsInfo.Id > 0 {
				players = append(players, rsInfo)
			}
		}
		ReturnSuccess(c, 0, "success", players, 1)
		return
	}
	rsDb, errDb := models.GetPlayers(aid, "score desc")
	if errDb == nil {
		for _, value := range rsDb {
			cache.Rdb.ZAdd(cache.Rctx, redisKey, cache.Zscore(value.Id, value.Score)).Err()
		}
		//遍历完成以后为rediskey设置过期时间
		cache.Rdb.Expire(cache.Rctx, redisKey, 24*time.Hour)
		ReturnSuccess(c, 0, "success", rsDb, 1)
		return
	}

	ReturnError(c, 4004, "没有相关信息")
	return
}
