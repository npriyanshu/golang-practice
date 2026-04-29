package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I am gonna learn switch case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("you got 1, so you can open")
	case 2:
		fmt.Println("you got 2, so you can eat")
	case 3:
		fmt.Println("you got 3, so you can sleep")
		// use fallthrough to execute the next case also
		fallthrough
	case 4:
		fmt.Println("you got 4, so you can code")
	case 5:
		fmt.Println("you got 5, so you can code")
	case 6:
		fmt.Println("you got 6, so you can code")
	default:
		fmt.Println("you got something else")
	}
}
