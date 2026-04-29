package main

import "fmt"

func main() {
	fmt.Println("Welcome to closures in Go")

	increment := counter();

	increment()
	increment()
	increment()

}

func counter() func () int {
	var x int = 0;

	return func () int {
		x++
		fmt.Println(x)
		return x
	}
}
