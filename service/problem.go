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

func TestHeader(c *gin.Context) {
	token := c.GetHeader("token")
	analyseToken, err := utils.AnalyseToken(token)
	if err != nil {

	}
	fmt.Println(analyseToken)
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
	categoryIdentities := c.PostFormArray("category_identities")
	testCases := c.PostFormArray("test_cases")

	//问题处理
	data, err := models.AddProblem(title, content, maxRuntime, maxMemory, categoryIdentities, testCases)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	//problemIdentity := data.(string)
	//
	////问题分类处理
	//for _, categoryIdentity := range categoryIdentities {
	//	_, err := models.AddProblemCategory(problemIdentity, categoryIdentity)
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"code": 400,
	//			"msg":  err.Error(),
	//			"data": nil,
	//		})
	//		return
	//	}
	//}
	//
	////问题测试用例处理
	//for _, testCase := range testCases {
	//	//testCase{"input":"1 2\n","output":"3"}
	//	testCaseMap := make(map[string]string, 0)
	//	json.Unmarshal([]byte(testCase), &testCaseMap)
	//	_, err := models.AddTestCase(problemIdentity, testCaseMap["input"], testCaseMap["output"])
	//	if err != nil {
	//		c.JSON(200, gin.H{
	//			"code": 400,
	//			"msg":  err.Error(),
	//			"data": nil,
	//		})
	//		return
	//	}
	//}
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
