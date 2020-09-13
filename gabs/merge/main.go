package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	obj1, _ := gabs.ParseJSON([]byte(`{"user":{"name":"dj"}}`))
	obj2, _ := gabs.ParseJSON([]byte(`{"user":{"age":18}}`))
	obj1.Merge(obj2)
	fmt.Println(obj1)

	arr1, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["game"]}}`))
	arr2, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	arr1.Merge(arr2)
	fmt.Println(arr1)

	obj3, _ := gabs.ParseJSON([]byte(`{"user":{"name":"dj", "hobbies": "game"}}`))
	arr3, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	obj3.Merge(arr3)
	fmt.Println(obj3)

	obj4, _ := gabs.ParseJSON([]byte(`{"user":{"name":"dj", "hobbies": "game"}}`))
	arr4, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	arr4.Merge(obj4)
	fmt.Println(arr4)

	obj5, _ := gabs.ParseJSON([]byte(`{"user":{"name":"dj", "hobbies": {"first": "game"}}}`))
	arr5, _ := gabs.ParseJSON([]byte(`{"user": {"hobbies": ["programming"]}}`))
	obj5.Merge(arr5)
	fmt.Println(obj5)
}
