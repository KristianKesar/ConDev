// Dining Philosophers
// Author: Joseph Kehoe
// Created: 21/10/24
// Modified by: Kristian Kesar
// Modified On: 05/10/2026
// License: GPL 3.0
//
// Five philosophers sit at a round table with one fork between each pair.
// To eat, a philosopher needs both forks beside them. If everyone grabs
// one fork at the same time, nobody can get a second one and they all
// wait forever (a deadlock). This version avoids that by always picking
// up the lower numbered fork first.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// think pauses for a random 0-4 seconds, then prints a message.
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // random number from 0 to 4
	time.Sleep(X * time.Second)
	fmt.Println("Phil: ", index, "was thinking")
}

// eat works the same as think. The philosopher must hold both forks first.
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5)) // random number from 0 to 4
	time.Sleep(X * time.Second)
	fmt.Println("Phil: ", index, "was eating")
}

// getForks picks up the two forks beside a philosopher.
// Each fork is a channel that holds one value: putting a value in means
// "I have this fork". If it is already full, the philosopher waits.
func getForks(index int, forks map[int]chan bool) {
	// %5 makes the table circular: philosopher 4's second fork is fork 0.
	first, second := index, (index+1)%5

	// Deadlock fix: always take the lower numbered fork first.
	// Only philosopher 4 is affected (takes 0 then 4 instead of 4 then 0).
	if first > second {
		first, second = second, first
	}

	forks[first] <- true  // pick up first fork (waits if taken)
	forks[second] <- true // pick up second fork (waits if taken)
}

// putForks puts both forks down by emptying their channels,
// so a waiting neighbour can pick them up.
func putForks(index int, forks map[int]chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
}

// doPhilStuff is one philosopher's life: think, get forks, eat, put forks
// back, repeated forever. Each philosopher runs this in its own goroutine.
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
	wg.Done() // never reached because the loop above is infinite
}

func main() {
	// The WaitGroup stops main from exiting before the philosophers run.
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount)

	// Create one fork (a channel with room for 1 value) per philosopher.
	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1)
	}

	// "go" starts each philosopher in the background so all 5 run at once.
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	}

	wg.Wait() // wait here forever; stop the program with Ctrl+C
}
