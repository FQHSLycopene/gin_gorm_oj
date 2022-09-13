package define

import (
	"gin_gorm_oj/utils"
	"github.com/gin-gonic/gin"
)

const (
	DefaultPage = "1"
	DefaultSize = "20"
)

func GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		analyseToken, err := utils.AnalyseToken(token)
		if err != nil {

		}
		c.Set("UserName", analyseToken.UserName)
		c.Set("UserIdenetity", analyseToken.Identity)
	}

}
