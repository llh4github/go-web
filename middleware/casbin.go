package middleware

import (
	"fmt"
	"strings"

	"gitee.com/llh-gitee/go-web/config"
	"github.com/gin-gonic/gin"
)

// 获取url
//
// 去掉配置中的url前缀。如果没有前缀则原样返回
func getURL(raw string) (url string) {
	prefix := config.GetURLPrefix()
	has := strings.HasPrefix(raw, prefix)
	if has {
		temps := strings.Split(raw, prefix)
		url = temps[1]
	} else {
		url = raw
	}
	return
}

// CasbinMiddleWare 权限认证中间件
// TODO 没搞明白，暂时先注释掉
func CasbinMiddleWare(c *gin.Context) {
	p := c.Request.URL.Path
	// m := c.Request.Method
	prefix := config.GetURLPrefix()
	// token := c.GetHeader("token")
	has := strings.HasPrefix(p, prefix)
	fmt.Printf("test:  %s , %v \n", p, has)
	/**
	var userName string
	userName = c.GetHeader("userName")
	if userName == "" {
		fmt.Println("headers invalid")
		c.JSON(200, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	// 请求的path
	// p := c.Request.URL.Path
	// // 请求的方法
	// m := c.Request.Method
	// // 这里认证
	// res, err := config.Enforcer
	// 这个 HasPermissionForUser 跟上面的有什么区别
	// EnforceSafe 会验证角色的相关的权限
	// 而 HasPermissionForUser 只验证用户是否有权限
	//res = Enforcer.HasPermissionForUser(userName,p,m)


	if err != nil {
		fmt.Println("no permission ")
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	if !res {
		fmt.Println("permission check failed")
		c.JSON(200, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    "",
		})
		c.Abort()
		return
	}
	*/
	c.Next()

}
