package main

import (
	"fmt"
	"time"
)

const (
	MAX = 10
)

type Leaky_Bucket struct {
	buffer chan int
}

func New_Leaky_Bucket() *Leaky_Bucket {
	lb := Leaky_Bucket{}
	lb.buffer = make(chan int, MAX)
	return &lb
}

func (lb *Leaky_Bucket) add(val int) {
	select {
	case lb.buffer <- val:
		fmt.Printf("Added %d to the buffer \n", val)
	default:
		fmt.Printf("Buffer full. Discarding %d\n", val)
	}
}

func (lb *Leaky_Bucket) run() {
	for {
		val := <-(lb.buffer)
		fmt.Printf("Removed %d from buffer ", val)
		fmt.Println(time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	lb := New_Leaky_Bucket()

	go lb.run()

	for i := 0; i < 20; i++ {
		i := i
		go lb.add(i)
	}

	time.Sleep(5 * time.Second)

}
