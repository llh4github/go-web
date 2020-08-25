package api

import (
	"gitee.com/llh-gitee/go-web/common"
	"gitee.com/llh-gitee/go-web/model"
	"gitee.com/llh-gitee/go-web/service"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

type user struct {
	baseAPI
}

func (m *user) Add(c *gin.Context) {
	service := service.User{}
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
		common.ExceptionByCode(common.DataBindError)
	}
	r := service.Add(user)

	m.respJSON(c, common.OkResponse(r))
}
func (m *user) login(c *gin.Context) {
	service := service.User{}
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
		common.ExceptionByCode(common.DataBindError)
	}
	rs := service.Login(user)
	m.respJSON(c, common.OkResponse(rs))
}
