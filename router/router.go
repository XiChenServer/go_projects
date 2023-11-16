package router

import (
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/gorilla/sessions"
	"go_ranking/config"
	"go_ranking/controllers"
	"go_ranking/pkg/logger"
)

func Router() *gin.Engine {
	r := gin.Default() //生成实例
	//controllers.InitSqlTable()
	//创建请求
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	r.Use(logger.Recover)
	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
	}
	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
	}
	vote := r.Group("/vote")
	{
		vote.POST("/add", controllers.VoteController{}.AddVote)
	}
	r.POST("/ranking", controllers.PlayerController{}.GetRanking)
	return r
}
