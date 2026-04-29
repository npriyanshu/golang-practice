package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")

	languages := make(map[string]string) // make a map of type string to string

	languages["Js"] = "javascript"
	languages["Py"] = "python"
	languages["Go"] = "golang"
	languages["Rs"] = "rust"

	fmt.Println("List of all languages", languages)

	fmt.Println("Js shorts for:",languages["Js"])

	// delete a key-value from a map 
	delete(languages, "Rs")

	fmt.Println("List of all languages", languages)


	// how to loop over maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n",key, value)
	}

	value,isPresent := languages["Rs"]

	if isPresent {
		fmt.Println("Value of Rust is", value)
	} else {
		fmt.Println("Value of Rust is not present")
	}

}
