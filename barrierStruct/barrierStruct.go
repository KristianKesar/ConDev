// Author: Kristian Kesar
// Date 28/09/2026
// People that have helped me:
// People I have helped:
// issue: make the barrier reusable (completed the barrier goes through 4 passes)
// license: GPL-3.0

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// Create a barrier data type
type barrier struct {
	Chan1   chan bool //change the single channel into 2 chans?
	Chan2   chan bool // first waits for all to arrive the second for all to leave
	theLock sync.Mutex
	total   int
	count   int
}

// creates a properly initialized barrier
// N== number of threads (go Routines)
func createBarrier(N int) barrier {
	theBarrier := barrier{
		Chan1: make(chan bool), //channel that counts arriving
		Chan2: make(chan bool), // channel for the ones leaving
		total: N,
		count: 0,
	}
	return theBarrier
}

// Method belonging to barrier data type
// Blocks until everyone reaches this point then lets everyone continue
func (b *barrier) wait() {
	b.theLock.Lock()
	b.count++
	if b.count == b.total {
		b.theLock.Unlock()  //unlock so the numbers can enter
		for _ = range b.total - 1 { //add numbers in decreasing order
			<-b.Chan1 // put them into channel 1
		}
	} else {
		fmt.Println(b.count) //print the count if the count doesn't equal total
		b.theLock.Unlock() // unlock the mutex
		b.Chan1 <- true //make chan1 appear full
	}
	// copy of code block above so that they could leave
	b.theLock.Lock()
	b.count-- //decrease counter
	if b.count == 0 { // down to 0
		b.theLock.Unlock()
		for _ = range b.total - 1 { // start removing the numbers out so the barrier is empty and ready to be reused
			<-b.Chan2
		}
	} else {
		fmt.Println(b.count)
		b.theLock.Unlock()
		b.Chan2 <- true
	}
} //wait

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theBarrier *barrier) bool {
	for pass := range 4 { //for loop for all the passes/reuses of the barrier
		var X time.Duration
		X = time.Duration(rand.IntN(5)) // random time in between printing
		time.Sleep(X * time.Second) //wait random time amount
		fmt.Println("Part A", Num, "pass", pass)
		//Rendezvous here
		theBarrier.wait() //wait until it is reusable and ready top run part b
		fmt.Println("Part B", Num, "pass", pass)
	}
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := createBarrier(5)
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &barrier)
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
