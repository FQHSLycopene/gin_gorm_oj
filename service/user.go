package service

import (
	"gin_gorm_oj/models"
	"gin_gorm_oj/utils"
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

// Login
// @Tags 公共方法
// @Summary 	登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json{"code":"200","msg":"","data",""}
// @Router /Login [Post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "必填信息为空",
			"data": nil,
		})
		return
	}
	//md5
	password = utils.GetMd5(password)
	data, err := models.Login(username, password)
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

// SendCode
// @Tags 公共方法
// @Summary 	发送验证码
// @Param Email formData string true "Email"
// @Success 200 {string} json{"code":"200","msg":"","data",""}
// @Router /SendCode [Post]
func SendCode(c *gin.Context) {
	Email := c.PostForm("Email")
	if Email == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "邮箱不能为空",
			"data": nil,
		})
		return
	}
	err := utils.SendCode("123", Email)
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
		"data": nil,
	})
}
