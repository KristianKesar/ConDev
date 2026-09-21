//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Kristian Kesar 21/09/2026
// Issues:
// The barrier is not implemented!
// People that have helped me: Filip Raguz, Tomas Radulescu
// People I have helped: Filip Raguz, Tomas Radulescu
// license: GPL-3.0
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, theLock *sync.Mutex, sem *semaphore.Weighted, ctx context.Context, count *int, total int) bool { // added the shared lock, semaphore, context, counter and total so every go routine uses the same ones
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	theLock.Lock() // only one go routine can change count at a time, otherwise two could read the same value and an increment would be lost
	*count++
	if *count == total {
		sem.Release(int64(total)) // give back all the tokens main took, so every waiting go routine can get one
	}
	theLock.Unlock()    // let the next go routine update count
	sem.Acquire(ctx, 1) // blocks while main holds all the tokens; only succeeds after the last go routine releases them
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	count := 0 // how many go routines have finished part A
	theLock.Lock()
	sem.Acquire(ctx, int64(totalRoutines)) // take every token so the semaphore starts at 0 and all go routines block at the barrier
	for i := range totalRoutines {
		go doStuff(i, &wg, &theLock, sem, ctx, &count, totalRoutines) // pass the shared lock, semaphore, context, counter and total
	}
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
