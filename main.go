package main

import (
	"demo4casbin/api"
	"github.com/sirupsen/logrus"
)

func main() {
	err := api.Router.Run()
	if err != nil {
		logrus.Error("gin 启动失败！")
	}
}
