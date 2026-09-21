// Author: Kristian Kesar
// Date 21/09/2026
// People that have helped me: Filip Raguz, Tomas Radulescu
// People I have helped: Filip Raguz, Tomas Radulescu
// license: GPL-3.0

//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Kristian Kesar
// Issues:
// None - barrier implemented with a mutex and a semaphore
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
	//wait here until everyone has completed part A
	theLock.Lock() // only one go routine can change count at a time
	*count++
	if *count == total { // true for exactly one go routine: the last one to finish part A
		sem.Release(int64(total))
	}
	theLock.Unlock() // let the next go routine update count
	sem.Acquire(ctx, 1)
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
	for i := range totalRoutines {         //create the go Routines here
		go doStuff(i, &wg, &theLock, sem, ctx, &count, totalRoutines)
	}
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
