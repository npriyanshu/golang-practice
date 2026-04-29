package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{}
	var fList = [4]string{}

	// uninitialized slice is nil and has length 0 and capacity 0
	var nums []int
	fmt.Println("nums is nil :",nums == nil)

	// second way to declare a slice is to use make function


	var nums2 = make([]int, 2,5) // 2 here is the initial length and 5 is the capacity of the slice and it will be initialized with zero values of int type which is 0

	// initialized slice is not nil and has length 2 and capacity 5 
	fmt.Println("nums2 is nil :",nums2 == nil, "and its length is", len(nums2), "and its capacity is", cap(nums2))
	fmt.Println("nums2 :",nums2)

	dList := append(fruitList, "Apple", "Peach", "Banana", "Mango", "Orange")
	fmt.Printf("Type of fruitlist is %T and type of fList is %T \n", fruitList, fList)

	fmt.Println("Fruit list is", fruitList)

	fmt.Println("d list is", dList, "and its length is", len(dList))

	// fruitList = append(dList[:3]) // sllicing from 1 to 3 index but index 3 is not included it starts from index 1 and ends at index 2

	fmt.Println("Fruit list is", fruitList)

	highScore := make([]int, 4) // make a slice of length 4 and type int

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore[4] = 345  // will throw an error

	fmt.Println("High score is", highScore)

	// but if we use append method now it will not throw an error
	highScore = append(highScore, 555, 666, 777)

	fmt.Println(highScore)

	// sort package
 	fmt.Println("are high score sorted",sort.IntsAreSorted(highScore))

	sort.Ints(highScore)
	fmt.Println("are high score sorted",sort.IntsAreSorted(highScore))
	fmt.Println(highScore)

	var courses = []string{"reactjs", "javascript", "swift", "python"}

	fmt.Println("List of courses is", courses)

	var index int = 2;

	courses = append(courses[:index],courses[index+1:]...)

	fmt.Println("List of courses is", courses)

	courses = append(courses[:0],courses[1:]...)

	fmt.Println("List of courses is", courses)

}
