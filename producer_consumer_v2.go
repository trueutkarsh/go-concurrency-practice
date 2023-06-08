package main

import (
	"fmt"
	"time"
)

// fixed buffer size

const MAX_SIZE = 5
const NUM_PRODUCERS = 10
const NUM_CONSUMERS = 5

// consumer -> remove from buffer

func consumer(buffer *chan int) {
	for val := range *buffer {
		fmt.Printf("Removed elem %d from buffer \n", val)
	}
}

// producer -> add to buffer

func producer(buffer *chan int, val int) {
	*buffer <- val
	fmt.Printf("Added elem %d to buffer \n", val)
}

func main() {

	buffer := make(chan int, MAX_SIZE)

	for i := 0; i < NUM_CONSUMERS; i++ {
		go consumer(&buffer)
	}

	for i := 0; i < NUM_PRODUCERS; i++ {
		go producer(&buffer, i)
	}

	time.Sleep(10 * time.Second)

}
