package main

import (
	"encoding/json"
	"fmt"
)

type JsonStruct struct {
	Array   []int       `json:"array"`
	Boolean bool        `json:"boolean"`
	Null    interface{} `json:"null"`
	Number  int         `json:"number"`
	Object  ObjectMap   `json:"object"`
	String  string      `json:"string"`
}

type ObjectMap struct {
	A string `json:"a"`
	C string `json:"c"`
	E string `json:"e"`
}

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
	var jsonStruct JsonStruct
	json.Unmarshal([]byte(jsonString), &jsonStruct)

	fmt.Println(fmt.Sprintf("%+v", jsonStruct))

	fmt.Println(jsonStruct.Object.A)
}
