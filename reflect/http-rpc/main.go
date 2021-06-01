package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type StringObject struct{}

func (StringObject) Concat(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

func (StringObject) ToUpper(s string) (string, error) {
	return strings.ToUpper(s), nil
}

func (StringObject) ToLower(s string) (string, error) {
	return strings.ToLower(s), nil
}

type MathObject struct{}

func (MathObject) Add(a, b int) (int, error) {
	return a + b, nil
}

func (MathObject) Sub(a, b int) (int, error) {
	return a - b, nil
}

func (MathObject) Mul(a, b int) (int, error) {
	return a * b, nil
}

func (MathObject) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divided by zero")
	}
	return a / b, nil
}

type RpcMethod struct {
	method reflect.Value
	args   []reflect.Type
}

var (
	mapObjMethods map[string]map[string]RpcMethod
)

func init() {
	mapObjMethods = make(map[string]map[string]RpcMethod)
}

func registerMethods(objName string, o interface{}) {
	v := reflect.ValueOf(o)

	mapObjMethods[objName] = make(map[string]RpcMethod)

	for i := 0; i < v.NumMethod(); i++ {
		m := v.Method(i)

		if m.Type().NumOut() != 2 {
			// 排除不是两个返回值的
			continue
		}

		if m.Type().Out(1).Name() != "error" {
			// 排除第二个返回值不是 error 的
			continue
		}

		t := v.Type().Method(i)
		methodName := t.Name
		if len(methodName) <= 1 || strings.ToUpper(methodName[0:1]) != methodName[0:1] {
			// 排除非导出方法
			continue
		}

		types := make([]reflect.Type, 0, 1)
		for j := 0; j < m.Type().NumIn(); j++ {
			types = append(types, m.Type().In(j))
		}

		mapObjMethods[objName][methodName] = RpcMethod{
			m, types,
		}
	}
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	m := map[string]interface{}{
		"error": err.Error(),
	}
	data, _ := json.Marshal(m)
	w.Write(data)
}

func response(w http.ResponseWriter, v interface{}) {
	w.WriteHeader(http.StatusOK)

	m := map[string]interface{}{
		"data": v,
	}
	data, _ := json.Marshal(m)
	w.Write(data)
}

func lookupMethod(name, method string) RpcMethod {
	m := mapObjMethods[name]
	if len(m) == 0 {
		return RpcMethod{}
	}

	return m[method]
}

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path[1:], "/")
	if len(parts) < 2 {
		handleError(w, errors.New("invalid request"))
		return
	}

	m := lookupMethod(parts[0], parts[1])
	if m.method.IsZero() {
		handleError(w, fmt.Errorf("no such method:%s in object:%s", parts[0], parts[1]))
		return
	}

	argSs := parts[2:]
	if len(m.args) != len(argSs) {
		handleError(w, errors.New("inconsistant args num"))
		return
	}

	argVs := make([]reflect.Value, 0, 1)
	for i, t := range m.args {
		switch t.Kind() {
		case reflect.Int:
			value, _ := strconv.Atoi(argSs[i])
			argVs = append(argVs, reflect.ValueOf(value))

		case reflect.String:
			argVs = append(argVs, reflect.ValueOf(argSs[i]))

		default:
			handleError(w, fmt.Errorf("invalid arg type:%s", t.Kind()))
			return
		}
	}

	ret := m.method.Call(argVs)
	err := ret[1].Interface()
	if err != nil {
		handleError(w, err.(error))
		return
	}

	response(w, ret[0].Interface())
}

func main() {
	registerMethods("math", MathObject{})
	registerMethods("string", StringObject{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
