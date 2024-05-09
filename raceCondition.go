// TODO
// Write two goroutines which have a race condition when executed concurrently.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Declare a shared variable
	var counter int

	// Launch two goroutines that increment the counter concurrently
	go increment("Goroutine 1", &counter)
	go increment("Goroutine 2", &counter)

	// Sleep for a while to allow goroutines to finish execution
	time.Sleep(2 * time.Second)

	// Print the final value of the counter
	fmt.Println("Final Counter Value:", counter)
}

func increment(name string, counter *int) {
	for i := 0; i < 5; i++ {
		// Simulate some computation
		time.Sleep(time.Millisecond * 100)

		// Increment the counter
		*counter++
		fmt.Printf("%s: Counter = %d\n", name, *counter)
	}
}
