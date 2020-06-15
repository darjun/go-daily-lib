package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	text := `
	[request_definition]
	r = sub, obj, act
  
	[policy_definition]
	p = sub, obj, act
  
	[role_definition]
	g = _, _
  
	[policy_effect]
	e = some(where (p.eft == allow))
  
	[matchers]
	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	`

	m, _ := model.NewModelFromString(text)
	a := fileadapter.NewAdapter("./policy.csv")
	e, _ := casbin.NewEnforcer(m, a)

	check(e, "dajun", "data1", "read")
	check(e, "lizi", "data2", "write")
	check(e, "dajun", "data1", "write")
	check(e, "dajun", "data2", "read")
}
