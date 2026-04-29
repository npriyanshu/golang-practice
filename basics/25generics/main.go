package main

import "fmt"

// print items of any type using generics

// func printItems[T any](items []T) // using any is not recommended because it does not provide any type safety and we can pass any type of slice to the function which can lead to runtime errors

// func printItems[T int | string | bool](items []T) { // union types are a better way to provide type safety and we can specify the types that we want to allow for the function
func printItems[T comparable](items []T) { // we can also use the comparable constraint to allow any type that can be compared using the == operator

	for _, item := range items {
		fmt.Println(item)
	}
}

// we can also use them in structs

type stack[T any] struct{
	elements []T
}

func main() {
	items := []int{1, 2, 3, 4, 5}
	printItems(items)

	myStack := stack[int]{
		elements: []int{1, 2, 3},
	}

	fmt.Println(myStack)
}
