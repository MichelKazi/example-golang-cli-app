package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
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
  "string": "ዩኒኮድ ወረጘ የጝ?"
}`

func main() {
	var jsonStruct JsonStruct

	stringSlice := []string{"dahlia", "toad", "phil", "doro", "doro"}

	stringSlice = append(stringSlice, "jeffry", "test")

	json.Unmarshal([]byte(jsonString), &jsonStruct)

	fmt.Print(len(jsonString), unsafe.Sizeof(jsonString))

	for i, v := range stringSlice {
		fmt.Println(i, v)
	}
}
