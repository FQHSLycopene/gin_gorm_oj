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
	r.GET("/Problem", service.GetProblemList)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/Problem/:id", service.GetProblemDetail)
	r.POST("/Problem", service.AddProblem)
	return r
}
