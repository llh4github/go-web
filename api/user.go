package api

import (
	"demo4casbin/common"
	"demo4casbin/model"
	"demo4casbin/service"
	"fmt"

	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

//  var service = &service.User{}

type user struct {
	baseAPI
}

func (m *user) Add(c *gin.Context) {
	service := service.User{}
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		log.Errorf("数据绑定错误, %v \n", err)
		m.respJSON(c, common.ErrorResponse(500, "数据绑定错误"))
	}
	fmt.Println("序列化后的user : ", user)
	r := service.Add(user)
	m.respJSON(c, common.OkResponse(r))
}
