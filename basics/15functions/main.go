package main

import "fmt"

func main() {

	fmt.Println("Welcome to functions in Go")

	fmt.Println(greet("Priyanshu"))

	fmt.Println("Enter the value of a");
	var a int
	fmt.Scan(&a)

	fmt.Println("Enter the value of b");
	var b int
	fmt.Scan(&b)

	fmt.Println("Enter the operator");
	var op string
	fmt.Scan(&op)

	fmt.Println(calculate(a,b,op))


	fmt.Println(proAdder(1,2,3,4,5,6))
}

func greet(name string) string {
	return "Hello " + name
}

func calculate(a int, b int, op string) int {
	switch op {
	case "+":
		return a + b

	case "-":
		return a - b

	case "*":
		return a * b

	default:
		return a / b
	}

}


// how to handle multiple values from a function
// variadic function
func proAdder(values ...int) int{
	total := 0;

	for _, v := range values {
		total += v;

	}
	return total;
}
