package main

import (
	"fmt"
	"reflect"
)

var (
	// name, course string
	// module       float64
	name   = "Nigel"            //Name of subscriber
	course = "Docker Deep Dive" //Name of course
	module = 3.2                //Current module
)

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
}