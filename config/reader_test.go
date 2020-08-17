package config

import "testing"

func TestReader(t *testing.T) {
	c := GetBDConfig()
	t.Log(c)
}
