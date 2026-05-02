package auth

import "fmt"


type User struct {
	Name string
	Email string
}

func Session() {
	fmt.Println("session created")
}


