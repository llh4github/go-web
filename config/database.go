package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // mysql dirve
	"github.com/jinzhu/gorm"
)

func init() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Printf("database connect error: \n %s \n", err)
	}
}
