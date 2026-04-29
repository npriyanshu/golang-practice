package main

import "fmt"

func main() {

	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	// fruitList[3] = "Banana"

	fmt.Println("Fruit list is", fruitList)

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Veg list is", vegList, " and the length is", len(vegList))

}