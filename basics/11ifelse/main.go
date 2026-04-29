package main

import "fmt"

func main() {

	fmt.Println("If else in golang")

	loginCount := 2
	var result string;

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 { 
    result = "Watch out"
	}else {
		result = "Premium User"
	}

	fmt.Println(result)


	if num :=loginCount; num > 10 {
		fmt.Println("number is greater than 10")
	} else {
		fmt.Println("number is less than 10")
	}
}
