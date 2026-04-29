package main

import (
	"fmt"
	"sync"
	"time"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("task", id, "done")
	fmt.Println("doing task", id)
}
func main() {
	var wg sync.WaitGroup
	for i := range 10 {
		wg.Add(1)
		go func(id int) {
			task(id, &wg)
		}(i)
	}
	fmt.Println("waiting for the tasks to be completed")
	wg.Wait()
	time.Sleep(time.Second * 1)

	fmt.Println("done")

}
