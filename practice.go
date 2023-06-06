package main

import (
	"fmt"
	"sync"
	"time"
)

// var MAX int = 5
// var NUM_PRODUCERS int = 5
const MAX_SIZE = 5
const NUM_PRODUCERS = 5
const NUM_CONSUMERS = 3

// lets create a semaphore example\

func worker(sem *chan int, id int) {
	// access sem
	// fmt.Println(fmt.Sprintf("%b", id))
	*sem <- id
	fmt.Println("Starting job ", id)
	time.Sleep(time.Millisecond)
	fmt.Println("Finished job", id)
	<-*sem
}

func main() {

	sem := make(chan int, 2)
	// done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i += 1 {
		i := i
		go func() {
			worker(&sem, i)
			wg.Done()
		}()
	}
	wg.Wait()
	// time.Sleep(10 * time.Millisecond)

}
