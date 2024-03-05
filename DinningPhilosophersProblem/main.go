package main

import (
	"fmt"
	"sync"
)

type ChopSticks struct{ sync.Mutex }
type Philosophers struct {
	number              int
	leftChop, rightChop *ChopSticks
}

var host = make(chan bool, 2)
var done = make(chan struct{})

func (p Philosophers) eat() {
	for j := 0; j < 3; j++ {
		host <- true

		p.leftChop.Lock()
		p.rightChop.Lock()

		fmt.Printf("starting to eat %d\n", j+1)
		fmt.Printf("finishing eating %d\n", j+1)

		p.rightChop.Unlock()
		p.leftChop.Unlock()

		<-host
	}
	done <- struct{}{}
}

func main() {
	var wg sync.WaitGroup
	chopSticks := make([]*ChopSticks, 5)
	for i := range chopSticks {
		chopSticks[i] = new(ChopSticks)
	}

	philosophers := make([]*Philosophers, 5)
	for i := range philosophers {
		philosophers[i] = &Philosophers{i + 1, chopSticks[i], chopSticks[(i+1)%5]}
	}

	for i := range philosophers {
		wg.Add(1)
		go philosophers[i].eat()
	}

	go func() {
		wg.Wait()
		close(done)
	}()
	<-done
	fmt.Println("All philosophers have finished eating three times.")
}
