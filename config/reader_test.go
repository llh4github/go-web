package config

import (
	"os"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	c := GetBDConfig()
	t.Log(c)
}

// 测试相对路径
func TestPath(t *testing.T) {
	dir, err := os.Getwd()

	if err != nil {
		t.Log("获取当前工作目录失败！", err)
	}
	t.Log(dir)
	paths := strings.Split(dir, "go-web")

	a, b := os.Stat(paths[0] + "/go-web")
	if b != nil {
		t.Error(b)
	}
	t.Log(a.IsDir())
}
