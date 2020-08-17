package security

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
)

// 模型结构体。临时用。
type modelTemp struct {
	Sub, // the user that wants to access a resource.
	Obj, // the resource that is going to be accessed.
	Act string // the operation that the user performs on the resource.
}

var enforce *casbin.Enforcer

func init() {

	var err error
	enforce, err = casbin.NewEnforcer("./_config/model.conf", "./_config/policy.csv")
	if err != nil {
		fmt.Println(err)
	}
}
func checkPerms(model modelTemp) {
	ok, err := enforce.Enforce(model.Sub, model.Obj, model.Act)
	if err != nil {
		fmt.Println(err)
	}
	if ok == true {
		fmt.Printf("%s has perm. \n", model.Sub)
	} else {
		fmt.Printf("%s has not perm. \n", model.Sub)
	}
}
func TestReadConf(t *testing.T) {

	m1 := modelTemp{
		Sub: "alice",
		Obj: "data1",
		Act: "read",
	}
	m2 := modelTemp{
		Sub: "Tom",
		Obj: "data1",
		Act: "read",
	}
	m3 := modelTemp{
		Sub: "bob",
		Obj: "data2",
		Act: "write",
	}
	checkPerms(m1)
	checkPerms(m2)
	checkPerms(m3)
}
