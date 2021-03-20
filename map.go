package main

import (
	"fmt"
)

func main() {
	leagueTItles := make(map[string]int)
	leagueTItles["Sunderland"] = 6
	leagueTItles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"Sunderland": 5,
		"Newcastle":  0,
	}
	fmt.Printf("\nLeague Titles: %v\nRecent head to heads: %v\n", leagueTItles, recentHead2Head)
}
