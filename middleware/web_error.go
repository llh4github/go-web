package middleware

import (
	"demo4casbin/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// HandleWebException 处理web异常
func HandleWebException(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Error("has error .")
			errorMsg(c, r)
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			c.Abort()
		}
	}()
	c.Next()
}

func errorMsg(c *gin.Context, r interface{}) {
	switch v := r.(type) {
	case error:
		otherErrorMsg(c, v.Error())
	case common.JSONWrapper:
		c.JSON(http.StatusOK, r)
	default:
		otherErrorMsg(c, r.(string))
	}
}

func otherErrorMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": common.UnknownError,
		"msg":  msg,
		"data": nil,
	})
}
