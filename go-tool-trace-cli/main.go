package main

import (
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// 1. Create the trace file
	f, _ := os.Create("trace.out")
	defer f.Close()

	// 2. Start Tracing
	trace.Start(f)
	defer trace.Stop()

	// 3. The "Simulated" Deadlock
	// We use a select with a timeout so we can exit and save the file.
	// In a real deadlock, this timeout wouldn't exist.
	ch := make(chan int)

	println("Attempting to send to channel (will block)...")

	select {
	case ch <- 1:
		// This will never happen because there is no receiver
	case <-time.After(100 * time.Millisecond):
		// This fires after 100ms, allowing us to see the "Gap"
		println("Timed out! Exiting to save trace.")
	}
}

