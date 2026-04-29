package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wait groups are counter that are used to wait for goroutines to finish
	// they are used to wait for a collection of goroutines to finish
	// this counter let main to wait for a goroutine to finish

	wg := sync.WaitGroup{} // create a wait group

	wg.Add(2) // add 2 to the counter

	go func(){
		defer wg.Done() // this will decrement the counter by 1 when the goroutine is done
       fmt.Println("task 1 started")
	   time.Sleep(150 * time.Millisecond)  
	   fmt.Println("task 1 finished")
	}()

	go func(){
		defer wg.Done() // this will decrement the counter by 1 when the goroutine is done
       fmt.Println("task 2 started")
	   time.Sleep(150 * time.Millisecond)  
	   fmt.Println("task 2 finished")
	}()

	// // if we add another task without incrementing the counter, the main function will not wait for it to finish and it will exit before the task is done
	// 	go func(){
	// 	defer wg.Done() // this will decrement the counter by 1 when the goroutine is done
    //    fmt.Println("task 3 started")
	//    time.Sleep(150 * time.Millisecond)  
	//    fmt.Println("task 3 finished")
	// }()

	wg.Wait() // this will block or wait for the counter to reach 0, which means that all goroutines have finished
	fmt.Println("all tasks finished")
}

