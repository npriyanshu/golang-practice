package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Welcome to web servers in golang")

}

func PerformGetRequest(){
	const myurl string = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}				
	defer response.Body.Close() // must close the res body after we are done with it
}