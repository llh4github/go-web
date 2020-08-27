package main

import (
	"gitee.com/llh-gitee/go-web/api"

	"github.com/sirupsen/logrus"
)

func main() {
	err := api.Router.Run(":8090")

	if err != nil {
		logrus.Error("gin 启动失败！")
	}
}
