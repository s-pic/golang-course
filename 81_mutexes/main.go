package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// solving the problem of the last lessen using a MuTex.
	// MuTex comes from mutual exclusion.
	// This means an explicit synchronization, where variable access is somewhat protected.

	// Prevent that reads and writes happen simoultaniously by using the two methods attached to
	// the Mutex type: lock() and unlock().
	// All code declared between a call to lock and unlock will happen in a single go routine.
	// Any other goroutines that want to access the locked variable will be blocked until the mutex is unlocked.
	// That avoids Race conditions.
	// Another way to solve the problem would be a channel.

	var m sync.Mutex

	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n--
			defer m.Unlock() // same effect like in first call
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("n is", n) // now 0. `go run -race main.go` does not log a race condition
}
