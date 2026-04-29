package main

import (
	"fmt"
	"time"
)

// we don't have constructors in go but we can hack around it and use a function as contructor

// as per convention use start the name of this function with New or new

func newAddress(state string, city string) *address {

	address := address{
		State : state,
		City : city,
	}
	return &address;
}

func newUser(name string, email string, status bool, age int) *User {
	user := User{
		Name : name,
		Email : email,
		Status : status,
		Age : age,
		CreatedAt: time.Now(),

	}
	return &user;
}

func main() {

	fmt.Println("Structs in golang")
	

	// priyanshu := User{"Priyanshu", "priyanshu@go.dev", true, 24, time.Now()} // struct literals

	// fmt.Println(priyanshu.CreatedAt.Format("01-02-2006 Monday 3:04:05 PM"))

	// // in the case of structure we can use %+v to get more detail

	// fmt.Printf("priyanshu details are: %+v\n", priyanshu)

	// fmt.Printf("Name is %v and email is %v.",priyanshu.Name,priyanshu.Email)

	// fmt.Println("updating the age with function")
	// priyanshu.updateUser(25)


	// fmt.Println(priyanshu)


	priyanshu := newUser("Priyanshu", "priyanshu@go.dev", true, 23)

	fmt.Println(priyanshu.Name,priyanshu.Age,priyanshu.Email,priyanshu.CreatedAt.Format("2006-01-02 Monday 03:04:05 PM"))


	// we can also create inline structs using struct literals

	language := struct{
		Name string
		Version int
	} {
		Name : "Go",
		Version : 1,
	}

	newAddress := newAddress("Maharashtra","Pune")

	// struct literal example
	anurag := User{
		Name : "anurag",
		Email : "anurag@go.dev",
		Status : true,
		Age : 23,
		CreatedAt: time.Now(),
		address: *newAddress,
	}

	language.Version = 2;

	anurag.City = "Patna"

	fmt.Println(language.Name,language.Version)
	fmt.Println(anurag)

}


// no inheritance in golang; no super or parent class
	// no constructor in golang

	// how to define a struct

	// we use type keyword to define a struct and then give it a name and fields

	// name is in capital letter because it is a struct

	type User struct {
		Name   string  // capital letters indicate public access
		Email  string
		Status bool
		Age    int
		CreatedAt time.Time // nanosecond precision
		address
	}

	type address struct {
		City string
		State string
	}


	// now its a method of user or structure
	// but if we want values in struct to actually change we have to make it pointer
	func (u *User) updateUser(v int) {
		// and struct automatically handles the dereferencing for us
		u.Age = v;
		fmt.Println("Age updated to :",u.Age)
	}