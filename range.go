package main

import (
	"fmt"
)

func main() {
	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Docker and Kubernetes"}
	coursesCompleted := []string{"Docker Deep Dive", "Go Fundamentals", "Puppet Fundamentals"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("The course " + i + "is in both lists")
			}
		}
	}

}
