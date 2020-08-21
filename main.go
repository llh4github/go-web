package main

import (
	"gitee.com/llh-gitee/go-web/api"
	"gitee.com/llh-gitee/go-web/config"

	"github.com/sirupsen/logrus"
)

func main() {
	err := api.Router.Run(":8090")
	defer config.MyDB.Close()
	if err != nil {
		logrus.Error("gin 启动失败！")

	}
}
