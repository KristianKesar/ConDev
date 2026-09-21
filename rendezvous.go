package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand/v2"
)
//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int) bool {
	var X time.Duration
	X=time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)//wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here

	fmt.Println("PartB",Num)
	wg.Done()
	return true
}



func main() {
	var wg sync.WaitGroup
	//barrier := make(chan bool)
	threadCount:=5

	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
