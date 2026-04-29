package main

import "fmt"

func main() {

	fmt.Println("learning loops")

	days :=  []string{"Sunday,	Monday,	Tuesday,	Wednesday,	Thursday,Friday,	Saturday"};

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i]);
	// }

	// better way

	for index := range days {
		fmt.Println(days[index])
	}


	// you can also get index and value

	for index, value := range days {
		fmt.Println("index:",index," value:",value)
	}


	nums := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};

	for index, value := range nums {
		fmt.Println("index:",index," value:",value)
	}

	// we can also use for loop like while loop

	var ifTen int = 1;

	for ifTen <= 10 {

		if ifTen == 5 {
			ifTen++
			continue
		}
		// break : will break the loop
		if ifTen == 8 {
			break
		}
		
		// continue : will skip the current iteration
		fmt.Println(ifTen)
		ifTen++
	}
	
}