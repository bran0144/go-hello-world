package main

import (
	"fmt"
)

func main() {

	switch "docker" {
	case "linux":
		fmt.Println("here are some recommended Linux courses")
	case "docker":
		fmt.Println("here are some recommended Docker courses")
	case "windows":
		fmt.Println("here are some recommended Windows courses")
	default:
		fmt.Println("sorry we couldn't find a match. Why not try our Top 100 list?")
	}
}
