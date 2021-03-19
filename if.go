package main

import (
	"fmt"
)

func main() {

	if firstRank, secondRank := 39, 614; firstRank < secondRank {
		fmt.Println("\nFirst course is doing better" + " than second course")
		if firstRank > 100 {
			fmt.Println("You may want to consider another job")
		}
	} else if firstRank > secondRank {
		fmt.Println("\nYour first course must be doing abysmally!")
		if secondRank > 100 {
			fmt.Println("You may wanna consider another job")
		}
	} else {
		fmt.Println("\nBoth courses are doing the same")
	}
}
