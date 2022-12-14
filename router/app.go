package router

import (
	_ "gin_gorm_oj/docs"
	"gin_gorm_oj/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//共有方法
	//问题
	r.GET("/Problem", service.GetProblemList)
	r.GET("/Problem/:identity", service.GetProblemDetail)
	//用户
	r.GET("/User/:identity", service.GetUser)
	r.POST("/Login", service.Login)
	r.POST("/SendCode", service.SendCode)
	r.POST("/Register", service.Register)
	//排行榜
	r.GET("/UserRank", service.GetUserRank)
	//提交记录
	r.GET("/Submit", service.GetSubmitList)
	//管理员私有方法
	r.POST("/Problem", service.AddProblem)

	//测试方法
	r.POST("/TestHeader", service.TestHeader)
	return r
}
