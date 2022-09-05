package service

import (
	"fmt"
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"gin_gorm_oj/utils"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json{"code":"200","msg":"","data",""}
// @Router /Problem [get]
func GetProblemList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	size := c.DefaultQuery("size", define.DefaultSize)
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	data, err := models.GetProblemList(page, size, keyword, categoryIdentity)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

func AddProblem(c *gin.Context) {
	data := models.ProblemBasic{}
	c.ShouldBind(&data)
	data.Identity = utils.GetUUID()
	problem := models.AddProblem(&data)
	c.JSON(200, gin.H{
		"data": problem,
	})
}

func GetProblemDetail(c *gin.Context) {
	identity := c.Param("identity")
	println(identity)
	data, err := models.GetProblemDetail(identity)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
