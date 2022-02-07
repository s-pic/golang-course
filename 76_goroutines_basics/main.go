package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
	}

	fmt.Println("f1 (goroutine) execution finished")
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")

	for i := 0; i < 3; i++ {
		fmt.Println("f2, i=", i)
	}

	fmt.Println("f2 (goroutine) execution finished")
}

func main() {
	fmt.Println("main execution started")

	fmt.Println("Nr of CPUs:", runtime.NumCPU())              // 16 cores
	fmt.Println("Nr of Go Routines:", runtime.NumGoroutine()) // 1 -> the main go routine

	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Architecture:", runtime.GOARCH)

	fmt.Println("That many Threads can be used to execute Go routines:",
		runtime.GOMAXPROCS(0), // just returns the current setting: 16 cores (the default value)
	)

	// launch goroutine (a function with an own stack, scheduled to an OS thread)

	go f1()                                                                 // now runs concurrently with main
	fmt.Println("Nr of Go Routines after go f1():", runtime.NumGoroutine()) // 1 -> the main go routine

	f2()

	// await f1 goroutine before finishing main (which stops everything)

	time.Sleep(time.Second * 2) // just to prove the concept, we would not do that in real world
	fmt.Println("main execution stopped")

}
