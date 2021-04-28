package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/casbin/casbin/v2"
)

func KeyMatch(key1, key2 string) bool {
	return strings.Index(key1, key2[:(len(key2)-1)]) == 0
}

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return (bool)(KeyMatch(name1, name2)), nil
}

func check(e *casbin.Enforcer, sub, obj, act string) {
	ok, _ := e.Enforce(sub, obj, act)
	if ok {
		fmt.Printf("%s CAN %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s CANNOT %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		log.Fatalf("NewEnforecer failed:%v\n", err)
	}

	e.AddFunction("my_func", KeyMatchFunc)

	check(e, "dajun", "data/1", "read")
	check(e, "dajun", "data/2", "read")
	check(e, "dajun", "data/1", "write")
	check(e, "dajun", "mydata", "read")
}
