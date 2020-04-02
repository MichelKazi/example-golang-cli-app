package main

import (
	"fmt"

	"github.com/protoworlock69/hello-world/parishiltonlongback"
)

type FakeStruct struct {
	Name     string
	Age      int
	Height   float64
	username string
	password string
}

func (f *FakeStruct) setPassword(password string) {
	f.password = password
}

func (f *FakeStruct) SetPassword(password string) {
	f.setPassword(password)
}

func main() {
	fmt.Println("Hello world")
	PrintHelloWorld()
	aFunction()
	fakeStruct := FakeStruct{Name: "Maggs"}
	fakeStruct.SetPassword("Who cares")
	parishiltonlongback.Longback()
}

func AFunction() {
	aFunction()
}

func aFunction() {
	fmt.Println("Private function")
}
