// package main

// func main(){

// 	// handline mutiple channels
// 	chan1 := make(chan string)
// 	chan2 := make(chan string)

// 	go func(){
// 		chan1 <- "ping";
// 	}()

// 	go func(){
// 		chan2 <- "pong";
// 	}()


// 	// we can use select statement to wait for multiple channels

// 	for i := 0; i< 2; i++ {
// 		select {
// 		case chan1val := <- chan1:
// 			println("value of chan 1", chan1val)
// 		case chan2val := <- chan2:
// 			println("value of chan 2", chan2val)
// 		}
// 	}
// }



package main

import (
	"fmt"
	"sync"
)


// we can also make our channels more type safe 
// by defining that a channel can only send or receive values of a specific type

// syntax will be chan<- type for send only channel and <-chan type for receive only channel in side the function parameters

// eg : task(id int, ch chan<- string) for a send only channel and task(id int, ch <-chan string) for a receive only channel
func task(id int , ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing task", id)
	ch <- fmt.Sprintf("task %d done", id) // send a message to the channel when the task is done
}

func main() {

	fmt.Println("Welcome to the channels")
	// channels are used to communicate between goroutines
	// they are used to send and receive values between goroutines
	// they are used to synchronize goroutines

ch := make(chan string) // create a channel of type string
var wg sync.WaitGroup

for i := range 10 {
	wg.Add(1)
		go task(i, ch, &wg)
	
}

// this goroutine will wait for all the tasks to be done and then close the channel

// we need to close the channel after all the tasks are done, otherwise the main function will be blocked forever waiting for messages from the channel

go func (){
	wg.Wait() // wait for all goroutines to finish
	close(ch) // close the channel after all goroutines have finished
}()

// receive messages from the channel
for i := range ch {
	fmt.Println(i)
}

}









// package main

// import (
// 	"fmt"
// 	// "sync"
// )

// func task(ch chan string, done chan bool) {
// 	defer func() {
// 		done <- true // signal that this task is done
// 	}()
// 	// defer wg.Done()
// 	for i := range ch {
// 		fmt.Println(i)
// 	}
// }

// func main() {

// 	fmt.Println("Welcome to the channels")
// 	// channels are used to communicate between goroutines
// 	// they are used to send and receive values between goroutines
// 	// they are used to synchronize goroutines

// 	ch := make(chan string) // create a channel of type string
// 	done := make(chan bool) // create a channel of type bool to signal when all tasks are done
// 	go task(ch, done)

// 	for i := range 10 {
// 		ch <- fmt.Sprintf("task %d done", i) // send a message to the channel when the task is done
// 	}

// 	// important to close the channel after all tasks are done, otherwise the goroutine will be waiting for more messages that will never come
// 	close(ch) // close the channel after all tasks are done
// 	<-done    // wait for the signal that all tasks are done

// }
