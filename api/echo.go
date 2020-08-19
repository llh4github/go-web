package api

import (
	"demo4casbin/common"

	"github.com/gin-gonic/gin"
)

type echoAPI struct {
	baseAPI
}

func (api echoAPI) hello(c *gin.Context) {
	api.respJSON(c, common.OkResponse("hello"))
}
