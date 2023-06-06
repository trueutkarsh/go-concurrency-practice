// To execute Go code, please declare a func main() in a package "main"
// # 1 bathroom, 3 people
// # republicans and democrats
// # 1. 3 people max
// # 2. cannot be together

// # democratUseBathroom(name)
// # republicanUseBathroom(name)

//

package main

import (
	"fmt"
	"sync"
	"time"
)

const MAX = 3

type Bathroom struct {
	num_dems int
	num_reps int

	lock sync.Mutex
	// empty		*sync.Cond
	full    *sync.Cond
	no_dems *sync.Cond
	no_reps *sync.Cond
}

func NewBathroom() *Bathroom {
	b := Bathroom{}
	b.full = sync.NewCond(&b.lock)
	b.no_dems = sync.NewCond(&b.lock)
	b.no_reps = sync.NewCond(&b.lock)
	return &b
}

func (b *Bathroom) democratUseBathroom(id int) {
	b.lock.Lock()
	for b.num_reps > 0 {
		b.no_reps.Wait()
	}

	for b.num_dems == MAX {
		b.full.Wait()
	}

	fmt.Printf("Democrate %d using bathroom \n", id)
	b.num_dems++
	// unlock
	b.lock.Unlock()
	time.Sleep(time.Second)
	b.lock.Lock()
	b.num_dems--
	b.full.Signal()

	if b.num_dems == 0 {
		b.no_dems.Signal()
		b.no_reps.Signal()
	}

	b.lock.Unlock()

}

func (b *Bathroom) republicanUseBathroom(id int) {
	b.lock.Lock()
	for b.num_dems > 0 {
		b.no_dems.Wait()
	}

	for b.num_reps == MAX {
		b.full.Wait()
	}

	fmt.Printf("Republican %d using bathroom \n", id)
	b.num_reps++
	// unlock
	b.lock.Unlock()
	time.Sleep(time.Second)
	b.lock.Lock()
	b.num_reps--
	b.full.Signal()

	if b.num_reps == 0 {
		b.no_reps.Signal()
		b.no_dems.Signal()
	}

	b.lock.Unlock()

}

func main() {

	numDemocrats := 5
	numRepublicans := 5

	var wg sync.WaitGroup

	b := NewBathroom()

	wg.Add(numDemocrats + numRepublicans)

	for i := 0; i < numRepublicans; i++ {
		i := i
		go func() {
			b.republicanUseBathroom(i)
			wg.Done()
		}()
	}
	for i := 0; i < numDemocrats; i++ {
		i := i
		go func() {
			b.democratUseBathroom(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
