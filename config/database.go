package config

import (
	"fmt"
)

func dbConfig() string {
	c := GetBDConfig()
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", c.Username, c.Password, c.Host, c.Dbname, c.Params)
	return s
}
