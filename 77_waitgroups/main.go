package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second) // simulate expensive task
	}

	fmt.Println("f1 (goroutine) execution finished")

	wg.Done() // notify wait group that this worker is done
	// same as (*wg).Done()
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f2, i=", i)
	}

	fmt.Println("f2 (goroutine) execution finished")
}

func main() {
	// same code as lesson 76

	// WAITGROUPS allow us to sync to block/wait until all routines
	// within a waitgroup have finished executing

	var wg sync.WaitGroup

	wg.Add(1) // the num of goroutines to wait for (1 -> f1)

	// launch goroutine (a function with an own stack, scheduled to an OS thread)

	go f1(&wg)                                                              // now runs concurrently with main
	fmt.Println("Nr of Go Routines after go f1():", runtime.NumGoroutine()) // 1 -> the main go routine

	f2()

	wg.Wait() // wait until all routines in the group succeeded
	fmt.Println("main execution stopped")

}
