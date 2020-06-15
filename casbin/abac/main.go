package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

type Object struct {
	Name  string
	Owner string
}

type Subject struct {
	Name string
	Hour int
}

func check(e *casbin.Enforcer, sub Subject, obj Object, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	} else {
		fmt.Printf("%s CANNOT %s %s at %d:00\n", sub.Name, act, obj.Name, sub.Hour)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	o := Object{"data", "dajun"}
	s1 := Subject{"dajun", 10}
	check(e, s1, o, "read")

	s2 := Subject{"lizi", 10}
	check(e, s2, o, "read")

	s3 := Subject{"dajun", 20}
	check(e, s3, o, "read")

	s4 := Subject{"lizi", 20}
	check(e, s4, o, "read")
}
