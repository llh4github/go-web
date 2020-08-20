package model

import (
	"testing"
)

func TestAdd(t *testing.T) {
	user := User{Username: "Tom"}
	user.SetCreatedInfo()
	saved := db.Create(&user)
	t.Log(saved)
}

func TestFindOne(t *testing.T) {
	user := User{}
	db.First(&user)
	t.Log(user)
	db.First(&user, 2) // where id = 2 LIMIT 1
	t.Log(user)
}
func TestFindWhere(t *testing.T) {
	users := make([]User, 10)
	db.Where("id is not null").Find(&users)
	t.Log(len(users))
}
