package service

import (
	"fmt"
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"strings"
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

// AddProblem
// @Tags 管理员私有方法
// @Summary 问题增加
// @Param token header string true "token"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData string true "max_runtime"
// @Param max_memory formData string true "max_memory"
// @Param category_identities formData array false "category_identities"
// @Param test_cases formData array true "test_cases"
// @Success 200 {string} json{"code":"200","msg":"","data",""}
// @Router /Problem [Post]
func AddProblem(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime := c.PostForm("max_runtime")
	maxMemory := c.PostForm("max_memory")
	categoryIdentitiesStr := c.PostForm("category_identities")
	//testCases := c.PostForm("test_cases")

	data, err := models.AddProblem(title, content, maxRuntime, maxMemory)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	problemIdentity := data.(string)
	categoryIdentities := strings.Split(categoryIdentitiesStr, ",")
	for _, categoryIdentity := range categoryIdentities {
		_, err := models.AddProblemCategory(problemIdentity, categoryIdentity)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"data": data,
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
