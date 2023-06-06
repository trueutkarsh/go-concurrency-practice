package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const MAX_SIZE = 5
const NUM_PRODUCERS = 10
const NUM_CONSUMERS = 5

func producer(buffer *[]int, guard *sync.Mutex, buffer_full *sync.Cond, buffer_empty *sync.Cond) {
	guard.Lock()
	// fmt.Printf("Producer %v \n", buffer)
	for len(*buffer) == MAX_SIZE {
		buffer_full.Wait()
	}
	elem := rand.Intn(100)
	*buffer = append(*buffer, elem)
	fmt.Printf("Added %d to buffer %v \n", elem, buffer)
	buffer_empty.Signal()
	guard.Unlock()
}

func consumer(buffer *[]int, guard *sync.Mutex, buffer_full *sync.Cond, buffer_empty *sync.Cond) {
	guard.Lock()
	// fmt.Printf("Consumer %v \n", buffer)
	for len(*buffer) == 0 {
		buffer_empty.Wait()
	}
	elem := (*buffer)[len(*buffer)-1]
	fmt.Printf("Removed elem %d from %v \n", elem, buffer)
	*buffer = (*buffer)[:len(*buffer)-1]
	buffer_full.Signal()
	guard.Unlock()
}

func main() {

	buffer := make([]int, 0)
	guard := sync.Mutex{}
	buffer_empty := sync.NewCond(&guard)
	buffer_full := sync.NewCond(&guard)

	for i := 0; i < NUM_CONSUMERS; i++ {

		go consumer(&buffer, &guard, buffer_full, buffer_empty)

	}

	for i := 0; i < NUM_PRODUCERS; i++ {

		go producer(&buffer, &guard, buffer_full, buffer_empty)

	}

	time.Sleep(10 * time.Second)
}
