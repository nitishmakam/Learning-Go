package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	leftCs, rightCs *Chopstick
	index           int
}

var c chan int // Used for permission
var on sync.Once
var wg sync.WaitGroup

func initialize() {
	c = make(chan int, 2) // Max of 2 philosophers can be eating at a time
}

func (p *Philosopher) Eat() {
	on.Do(initialize)
	for i := 0; i < 3; i++ {

		p.leftCs.Lock()
		p.rightCs.Lock()
		c <- p.index
		fmt.Println("starting to eat ", p.index)
		time.Sleep(20 * time.Millisecond)
		fmt.Println("finishing eating ", p.index)
		<-c
		p.leftCs.Unlock()
		p.rightCs.Unlock()

	}
	wg.Done()
}

func main() {
	chopSticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopSticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{chopSticks[i], chopSticks[(i+1)%5], i}
	}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philosophers[i].Eat()
	}
	wg.Wait()
}
