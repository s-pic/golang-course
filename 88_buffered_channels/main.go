package main

import (
	"fmt"
	"time"
)

func main() {
	bufferedChannel := make(chan int, 3) // capacity set at declaration

	// the sender of a buffered channel
	// will only block when there is no empty slot in the channel
	// while the receiver will block when its empty ??

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("goroutine starts. i is %d", i)
			c <- i // this blocks the goroutine until main wakes up
			fmt.Println("goroutine exits. i is %d", i)
		}
		close(bufferedChannel) // to tell the receiver that all data has been send.
		// We do not close all channels, but in this case that is important
	}(bufferedChannel)

	fmt.Println("main goroutine sleeps for 2 seconds") // so that other goroutines have time to start
	time.Sleep(time.Second * 2)

	for receivedVal := range bufferedChannel {
		fmt.Println("main goroutine received value from channel:", receivedVal)

		// we can send up to three values in the channel, before it starts blocking.
		// the receiver will block in the channel, when it is empty.

		// So the difference between a buffered an unbuffered channel:
		// Communication over an unbuffered channel casues the receiving and the sending
		// goroutines to synchronize. That's why they are called synchronous channels.

		// When using a buffered channel, we decouple sender and receiver.

		// Analogy to get the difference:
		// - Three cooks work in a kitchen.
		// - THe first one does the dough, the second one the baking and the thrird one the icing.
		// When using an unbuffered channel, each cook has to wait with passing the cake
		// onto the next cook until the receiving cook is ready to take the cake.
		// When using a buffered channel, we place little tables (staging areas) between the cooks

	}

}
