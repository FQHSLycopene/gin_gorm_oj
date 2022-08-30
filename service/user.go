package service

import (
	"gin_gorm_oj/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "id 不是 int类型",
			"data": nil,
		})
		return
	}
	data, err := models.GetUser(id)
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
