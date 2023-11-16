package models

import (
	"github.com/jinzhu/gorm"
	"go_ranking/dao"
)

type Player struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int    `json:"score"`
	//AddTime     int64  `json:"addTime"`
	//UpdateTime  int64  `json:"updateTime"`
}

func (Player) TableName() string {
	return "player"
}
func GetPlayers(aid int, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error
	return players, err
}

func GetPlayerInfo(id int) (Player, error) {
	var player Player
	err := dao.Db.Where("id = ?", id).First(&player).Error
	return player, err
}
func UpdatePlayerScore(id int) {
	var player Player
	dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
