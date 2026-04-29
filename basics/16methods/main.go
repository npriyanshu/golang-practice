package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")
	

	priyanshu := User{"Priyanshu", "priyanshu@go.dev", true, 24, 0} // struct literals

	fmt.Println(priyanshu)

	// in the case of structure we can use %+v to get more detail

	fmt.Printf("priyanshu details are: %+v\n", priyanshu)

	fmt.Printf("Name is %v and email is %v.",priyanshu.Name,priyanshu.Email)


	fmt.Println(" the status is :",priyanshu.GetStatus())

	priyanshu.newMail()

	
	fmt.Printf("Name is %v and email is %v.",priyanshu.Name,priyanshu.Email)




}


	type User struct {
		Name   string  // capital letters indicate public access
		Email  string
		Status bool
		Age    int
		oneAge int  // private not exported
	}

	// create a method by passing struct as a parameter

	func (u User) GetStatus() bool {
		return u.Status;
	}

	func (u User) newMail() {
		 u.Email = "test@go.dev"
		 fmt.Println("Email of this user is", u.Email)
	}