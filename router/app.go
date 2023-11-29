package router

import (
	_ "gin_gorm_oj/docs"
	"gin_gorm_oj/middlewares"
	"gin_gorm_oj/service"
	"gin_gorm_oj/test_api"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	//Swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	test := r.Group("/test")
	{
		test.GET("/get_test", test_api.TestApi{}.GetTest)
		test.POST("/post_test", test_api.TestApi{}.PostTest)
	}
	//问题
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)
	//用户
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)

	//排行榜
	r.GET("/rank-list", service.GetRankList)

	//提交记录
	r.GET("/submit-list", service.GetSubmitList)

	//管理员私有
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	{
		authAdmin.POST("/problem-create", service.ProblemCreate)
		//问题修改
		authAdmin.PUT("/problem-modify", service.ProblemModify)

		//分类列表
		authAdmin.GET("/category-list", service.GetCategoryList)
		//分类创建
		authAdmin.POST("/category-create", service.CategoryCreate)
		//分类删除
		authAdmin.DELETE("/category-delete", service.CategoryDelete)
		//分类修改
		authAdmin.PUT("/category-modify", service.CategoryModify)
	}
	//用户的私有方法
	authUser := r.Group("/user", middlewares.AuthUserCheck())
	{
		authUser.POST("/submit", service.Submit)
	}
	return r
}
