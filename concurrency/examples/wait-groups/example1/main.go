package main

import (
	"fmt"
	"sync"
	"time"
)

// Creating a global wait group.
var wg sync.WaitGroup

func worker(id int, timer time.Time) {
	// Mark the Goroutine as completed which will update the internal counter, and as soon as the
	// counter of the global wg (wait group) reaches zero, main goroutine is terminated.
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// pausing the Goroutine to mimic execution of worker.
	time.Sleep(time.Second)
	fmt.Printf("Worker %d is done at %s.\n", id, time.Since(timer))
}

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increments the counter to indicate that a Goroutine has been created.
		start := time.Now()
		go worker(i, start)
	}
	wg.Wait() // Waits until the internal counter becomes zero.
}
