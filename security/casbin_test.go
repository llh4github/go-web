package security

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin/v2"
)

var enforce *casbin.Enforcer

func init() {

	var err error
	enforce, err = casbin.NewEnforcer("./_config/model.conf", "./_config/policy.csv")
	if err != nil {
		fmt.Println(err)
	}
}

func TestReadConf(t *testing.T) {
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := enforce.Enforce(sub, obj, act)
	if err != nil {
		t.Log(err)
	}
	if ok == true {
		t.Logf("%s has perm.", sub)
	} else {
		t.Logf("%s has not perm.", sub)
	}
}
