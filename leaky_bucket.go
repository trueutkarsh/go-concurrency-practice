package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MAX = 10
)

type Leaky_Bucket struct {
	mut    sync.Mutex
	buffer []int
}

func (lb *Leaky_Bucket) add(val int) {
	lb.mut.Lock()
	if len(lb.buffer) < MAX {
		lb.buffer = append(lb.buffer, val)
		fmt.Printf("Len %d at %v\n", len((*lb).buffer), time.Now())
	}
	lb.mut.Unlock()
}

func (lb *Leaky_Bucket) run() {
	for true {
		lb.mut.Lock()
		if len(lb.buffer) > 0 {
			val := lb.buffer[0]
			lb.buffer = lb.buffer[1:]
			fmt.Printf("Len %d Processed %d at %v\n", len((*lb).buffer), val, time.Now())
		}
		lb.mut.Unlock()
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {

	lb := Leaky_Bucket{}

	go lb.run()

	// lb.run()
	// lb.run()
	for i := 0; i < 200; i++ {
		i := i
		go func() {
			time.Sleep(time.Duration(300+rand.Intn(1000)) * time.Millisecond)
			lb.add(i)
		}()
	}

	time.Sleep(5 * time.Second)

}
