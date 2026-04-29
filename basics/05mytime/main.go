package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("Welcome to time study")
	fmt.Println(presentTime.Format("01-02-2006 Monday 3:04:05 PM"))

	createdDate := time.Date(2026, time.June, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday 3:04:05 PM"))
}
