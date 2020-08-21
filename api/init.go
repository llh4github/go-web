package api

import (
	"demo4casbin/common"

	"github.com/gin-gonic/gin"
)

// Router 全局gin实例
var Router *gin.Engine

// APIGroup 当前使用的API组
var APIGroup *gin.RouterGroup

func init() {
	Router = gin.Default()
	APIGroup = Router.Group("api")
	APIGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "hello"})
	})
	registerEcho()
	registerUser()
}

// 基础API结构体。
// 虽然可以不使用结构体，但为了使用公共方法和避免函数名冲突，
// 所以选用使用结构体来管理各类API处理方法。
type baseAPI struct {
}

func (baseAPI) respJSON(c *gin.Context, data common.JSONWrapper) {
	c.JSON(200, data)
}

func registerEcho() {
	api := echoAPI{}
	APIGroup.GET("echo", api.hello)
	APIGroup.GET("echo/who", api.helloSomeone)
	APIGroup.POST("echo/girl", api.helloGirl)
}
func registerUser() {
	r := user{}
	APIGroup.POST("user", r.Add)
}
