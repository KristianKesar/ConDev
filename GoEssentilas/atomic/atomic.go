// Author: Kristian Kesar
// Date 21/09/2026
// People that have helped me: Filip Raguz, Tomas Radulescu
// People I have helped: Filip Raguz, Tomas Radulescu
// license: GPL-3.0
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Global variables shared between functions --A BAD IDEA
var wg sync.WaitGroup

func addsAtomic(n int, total *atomic.Int64) bool {
	for i := 0; i < n; i++ {
		total.Add(1)
	}
	wg.Done() //let waitgroup know we have finished
	return true
}

func main() {
	var total atomic.Int64
	//for loop using range option
	for i := range 10 {
		//the waitgroup is used as a barrier
		// init it to number of go routines
		wg.Add(1)
		fmt.Println("go Routine ", i)
		go addsAtomic(1000, &total)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done
	fmt.Println(total.Load())
}

//verified the outputis correct and working just like the other classmates
