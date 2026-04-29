package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int;
	// fmt.Println(ptr)

	myNumber:= 23;
	var ptr *int = &myNumber;

	var dptr **int = &ptr;

	fmt.Println("Value of myNumber is", myNumber)
	fmt.Println("Value of ptr and address of myNumber is", ptr)
	fmt.Println("Value of *ptr is", *ptr)
	fmt.Println("Value of dptr is",dptr)
	fmt.Println("Value at dptr is",*dptr)
	fmt.Println("Value at *dptr is",**dptr)

}