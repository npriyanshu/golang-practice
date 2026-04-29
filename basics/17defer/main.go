package main

import "fmt"

func main() {

	// defer follows the pattern of LIFO Last In First Out

	defer fmt.Println("World")
	defer fmt.Println("two")
	defer fmt.Println("one")

	fmt.Println("Hello")

	defer fmt.Print("\n")
	myDefer()
}


func myDefer(){

	for i := range 5 {
		defer fmt.Print(i)
	}
}