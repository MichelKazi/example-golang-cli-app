package main

import (
	"encoding/json"
	"fmt"
)

var jsonString = `{
  "array": [
	1,
	2,
	3
  ],
  "boolean": true,
  "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f"
  },
  "string": "Hello World"
}`

func main() {
	var jsonMap = make(map[string]interface{})
	json.Unmarshal([]byte(jsonString), &jsonMap)

	for key, value := range jsonMap {
		fmt.Println(key, value)
	}
}
