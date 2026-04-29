package main

import (
	"fmt"
	// "net/http"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/posts/1?id=5"

func main() {
	fmt.Println("Welcome to urls handling")

	fmt.Println("url :", myurl)

	// parsing the url

	result, _ :=url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of query params is: %T\n", qParams)

	fmt.Println(qParams["id"])

	// creating a url from scratch

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "jsonplaceholder.typicode.com",
		Path: "/posts/2",
		RawQuery: "id=5",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url is:", anotherUrl)


}