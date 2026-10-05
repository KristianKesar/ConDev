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
// Modified by: Kristian Kesar (c00296348@setu.ie)
// Modified on: 01/10/2026 and 05/10/2026
// Description: Make reusable barrier
// A simple barrier implemented using mutex and unbuffered channel
// license: GPL-3.0
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic" // to complete task 1, mutex to atomic
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *atomic.Int32, max int, wg *sync.WaitGroup, theChan chan bool, theChan2 chan bool) bool {
	for pass := range 3 { //loop to show the reusable passes
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("Part A", goNum, "pass", pass)
		//we wait here until everyone has completed part A
		if arrived.Add(1) == int32(max) { //changed to atomic values for the insertion
			for range max - 1 { // adding until all have arrived
				theChan <- true // send true to show channel is full
			}
		} else { //not all here yet we wait until signal
			<-theChan
		} //end of if-else
		if arrived.Add(-1) == 0 { // decrementing until it is 0
			for range max - 1 {
				theChan2 <- true // send true to show channel is empty
			}
		} else {
			<-theChan2
		}
		fmt.Println("PartB", goNum, "pass", pass)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	var arrived atomic.Int32 //arrived changed from regular int to atomic
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	theChan := make(chan bool) //use channel in place of semaphore
	theChan2 := make(chan bool) // channel 2 in place for the decrementing
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, theChan, theChan2) // inports for doStuff changed
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
