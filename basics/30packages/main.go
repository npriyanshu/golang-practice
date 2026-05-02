package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/npriyanshu/golangpractise/auth"
)

// to install a package use the command go get <package-name>
// for example := go get github.com/fatih/color

func main() {
	val := auth.Login("priyanshu", "secret")
	fmt.Println(val)
	auth.Session()

	user := auth.User{
		Name:  "Priyanshu",
		Email: "priyanshu@go.dev",
	}
	color.Green("User Name: %s, User Email: %s", user.Name, user.Email)
}
