package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var JSONrecord = `{
    "Flag": true,
    "Array": ["a","b","c"],
    "Entity": {
      "a1": "b1",
      "a2": "b2",
      "Value": -456,
      "Null": null
},
    "Message": "Hello Go!"
  }`

func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("Type is string:", k, c)
		case float64:
			fmt.Println("Type is float64:", k, c)
		case bool:
			fmt.Println("Type is bool:", k, c)
		case map[string]interface{}:
			fmt.Println("Type is a map:", k, c)
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("Type is %v: %T!\n", k, c)
		}
	}
	return
}

func exploreMap(m map[string]interface{}) {
	for k, v := range m {
		embeddedMap, ok := v.(map[string]interface{})
		if !ok {
			fmt.Printf("%v: %v\n", k, v)
		} else {
			fmt.Printf("{\"%v\": \n", k)
			exploreMap(embeddedMap)
			fmt.Printf("}\n")
		}

	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("using default json")
	} else {
		JSONrecord = os.Args[1]
	}
	JSONMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	exploreMap(JSONMap)
	typeSwitch(JSONMap)
}
