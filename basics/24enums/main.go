package main

import "fmt"

// enumerated types
// we can use enumerated types to define a set of named constants

// we can also create our custom types like this
// type MyType int
// we use custome types create enumerated types

// type OrderStatus int

// const (
// 	Received OrderStatus = iota // iota is a predeclared identifier that increment by 1 for each constant in the block
// 	Confirmed
// 	Shipped
// 	Delivered
// )

// we can also use the string type for enumerated types

type OrderStatus string

const (
	Received  OrderStatus = "received" // we can assign any string value to the constant
	Confirmed OrderStatus = "confirmed"
	Shipped   OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changed the order status to", status)
}
func main() {
	// changeOrderStatus(1) // we can pass the integer value of the constant to the function
	changeOrderStatus("received") // we can also pass the string value of the constant to the function
	changeOrderStatus(Delivered)  // we can also pass the constant itself to the function
}
