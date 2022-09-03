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

	//问题
	r.GET("/Problem", service.GetProblemList)
	r.GET("/Problem/:identity", service.GetProblemDetail)
	r.POST("/Problem", service.AddProblem)

	//用户
	r.GET("/User/:id", service.GetUser)
	r.POST("/Login", service.Login)
	r.POST("/SendCode", service.SendCode)
	r.POST("/Register", service.Register)

	//提交记录
	r.GET("/Submit", service.GetSubmitList)
	return r
}
