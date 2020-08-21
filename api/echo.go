package api

import (
	"fmt"
	"gitee.com/llh-gitee/go-web/common"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type echoAPI struct {
	baseAPI
}
type girl struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (api echoAPI) hello(c *gin.Context) {
	api.respJSON(c, common.OkResponse("hello"))
}
func (api echoAPI) helloSomeone(c *gin.Context) {
	name := c.Query("name")
	api.respJSON(c, common.OkResponse("hello "+name))
}

func (api echoAPI) helloGirl(c *gin.Context) {
	var g girl
	if err := c.ShouldBind(&g); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
		api.respJSON(c, common.ErrorResponse(500, "数据绑定错误"))
	}
	s := fmt.Sprintf("Hello girl , your name is %s , and age is %d", g.Name, g.Age)
	api.respJSON(c, common.OkResponse(s))
}
