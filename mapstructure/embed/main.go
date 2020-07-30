package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
}

type Friend1 struct {
	Person
}

type Friend2 struct {
	Person `mapstructure:",squash"`
}

func main() {
	datas := []string{`
		{ 
			"type": "friend1",
			"person": {
				"name":"dj"
			}
		}
	`,
		`
		{
			"type": "friend2",
			"name": "dj2"
		}
	`,
	}

	for _, data := range datas {
		var m map[string]interface{}
		err := json.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatal(err)
		}

		switch m["type"].(string) {
		case "friend1":
			var f1 Friend1
			mapstructure.Decode(m, &f1)
			fmt.Println("friend1", f1)

		case "friend2":
			var f2 Friend2
			mapstructure.Decode(m, &f2)
			fmt.Println("friend2", f2)
		}
	}
}
