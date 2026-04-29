package main

import (
	"fmt"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts/1"

func main() {

	fmt.Println("Web request")


	response, err :=http.Get(url);

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // it is important to close the response body after we are done with it, otherwise it can lead to memory leaks


	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)
}