package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
         reader := bufio.NewReader(os.Stdin);

	fmt.Println("What is your name?")
	name,_ := reader.ReadString('\n')

	var age int;
	var gender string;
	fmt.Println("What is your age?")

	n,err := fmt.Scan(&age,&gender)

	if(err != nil) {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Hello!",name,age,n)
	
	
}
