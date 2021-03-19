package main

import (
	"fmt"
	"os"
)

// var (
// 	// name, course string
// 	// // module       float64
// 	name   = "Nigel"            //Name of subscriber
// 	course = "Docker Deep Dive" //Name of course
// 	module = 3.2                //Current module
// )

func main() {
	name := os.Getenv("USERNAME")
	course := "Docker Deep Dive"

	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {
	*course = "First look: Native Docker Clustering"
	fmt.Println("\nTrying to change your course to", *course)
	return *course
}
