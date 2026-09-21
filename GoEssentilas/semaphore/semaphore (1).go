// Author: Kristian Kesar
// Date 21/09/2026
// People that have helped me: Filip Raguz, Darian Byrne, Tomas Radulescu
// People i have helped: Filip Raguz, Tomas Radulescu
// license: GPL-3.0
package main

import (
	"fmt"
	"sync"
	"time"
)

// make struct containing channel
// add init, acquire and release
type semaphore struct {
	theCounter chan struct{}
}

func Init(n int) *semaphore { // init keyword possibly? so im using capital I
	return &semaphore{make(chan struct{}, n)}
	// new function
}
func Acquire(sem *semaphore) {
	sem.theCounter <- struct{}{}
	// had to fix typos and add body
}
func Release(sem *semaphore) {
	<-sem.theCounter
	//made new function
}

func main() {
	maxGoroutines := 5
	sem := Init(maxGoroutines) //change line with the func that does the work

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Acquire(sem)       //change line with the func that does the work
			defer Release(sem) //change line with the func that does the work
			// Simulate a task
			fmt.Printf("Running task %d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	wg.Wait()
}

//verified the outputis correct and working just like the other classmates
