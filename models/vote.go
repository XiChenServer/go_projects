package models

type Vote struct {
	Id       int   `json:"id"`
	UserId   int   `json:"userId"`
	PlayerId int   `json:"playerId"`
	AddTime  int64 `json:"addTime"`
}
