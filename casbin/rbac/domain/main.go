package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func check(e *casbin.Enforcer, sub, domain, obj, act string) {
	ok, _ := e.Enforce(sub, domain, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s in %s\n", sub, act, obj, domain)
	} else {
		fmt.Printf("%s CANNOT %s %s in %s\n", sub, act, obj, domain)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	check(e, "dajun", "tenant1", "data1", "read")
	check(e, "dajun", "tenant2", "data2", "read")
}
