package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Hugo"
	age := 20
	version := 1.1
	fmt.Println("Hello, Mr.", name, "your age is", age)
	fmt.Println("This pragram is in version", version)

	fmt.Println("TypeOf name: ", reflect.TypeOf(name))
	fmt.Println("TypeOf age: ", reflect.TypeOf(age))
	fmt.Println("TypeOf version: ", reflect.TypeOf(version))
}
