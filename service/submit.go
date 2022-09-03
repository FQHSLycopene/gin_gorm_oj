package service

import (
	"gin_gorm_oj/define"
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 	提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_id query string false "problem_id"
// @Param user_id query string false "user_id"
// @Success 200 {string} json{"code":"200","msg":"","data",""}
// @Router /GetSubmitList [get]
func GetSubmitList(c *gin.Context) {
	page := c.DefaultQuery("page", define.DefaultPage)
	size := c.DefaultQuery("size", define.DefaultSize)
	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	data, err := models.GetSubmitList(page, size, problemIdentity, userIdentity)
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
