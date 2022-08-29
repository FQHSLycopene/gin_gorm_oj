package service

import (
	"fmt"
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @BasePath /api/v1

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_id query string false "category_id"
// @Success 200 {string} json{"code":"200","msg":"","data",""}
// @Router /GetProblemList [get]
func GetProblemList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	size := c.DefaultQuery("size", define.DefaultSize)
	keyword := c.Query("keyword")
	category_id := c.Query("category_id")
	data, err := models.GetProblemList(page, size, keyword, category_id)
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
	problem := models.AddProblem(&data)
	c.JSON(200, gin.H{
		"data": problem,
	})
}

func GetProblemDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "id 不是 int类型",
			"data": nil,
		})
	}
	data := models.GetProblemDetail(id)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}
