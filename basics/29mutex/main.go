package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu sync.Mutex // add a mutex to protect the views field 
}

func (p *post) inc(wg *sync.WaitGroup)() {
	defer wg.Done();  // must use defer to ensure that Done is called even if panic occurs
	p.mu.Lock() // lock the mutex before modifying the views field
	p.views++
	defer p.mu.Unlock() // unlock the mutex after modifying the views field and use defer to ensure that it is unlocked even if panic occurs

}

func main() {

	var wg sync.WaitGroup;

	myPost := post{views: 0}


	 for i:=0; i<100; i++{
		wg.Add(1);
	 go myPost.inc(&wg)
	 }

	 wg.Wait();


	 // for now the output could be 100, 99, 98, etc because of the race condition

	 // to solve this we need to use a mutex to protect the views field from concurrent access
	fmt.Println(myPost.views)

}