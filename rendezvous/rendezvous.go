// Author: Kristian Kesar
// Date 21/09/2026
// People that have helped me: Filip Raguz, Tomas Radulescu
// People I have helped: Filip Raguz, Tomas Radulescu
// license: GPL-3.0

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, theLock *sync.Mutex, count *int, total int, barrier chan bool) bool { // added the shared lock, counter, total and channel so every go routine uses the same ones
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait a random amount of time
	fmt.Println("Part A", Num)
	//Rendezvous here
	theLock.Lock() // only one go routine can change count at a time
	*count++
	if *count == total {
		close(barrier) // closing wakes every go routine waiting on the channel at once (a single send would only wake one)
	}
	theLock.Unlock() // let the next go routine update count
	<-barrier
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrier := make(chan bool)
	var theLock sync.Mutex // protects count
	count := 0
	threadCount := 5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, &theLock, &count, threadCount, barrier) // pass the shared lock, counter, total and channel
	}
	wg.Wait() //wait here until everyone (5 go routines) is done

}
