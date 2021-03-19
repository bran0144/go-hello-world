package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	switch tempNum := random(); tempNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got an even numter -", tempNum)
	case 1, 3, 5, 7, 9:
		fmt.Println("We got an odd numter -", tempNum)
	}
}
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
