package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// processTenant simulates work being done for a specific tenant.
// It can sometimes fail, returning an error.
func processTenant(tenantID string) (string, error) {
	log.Printf("Processing tenant with ID: %q", tenantID)
	// Simulate variable work duration
	time.Sleep(time.Millisecond * time.Duration(100+len(tenantID)*10))

	// Simulate a potential failure for demonstration
	if tenantID == "tenant-2" {
		return "", fmt.Errorf("failed to process tenant: %s", tenantID)
	}

	result := fmt.Sprintf("Successfully processed tenant %s", tenantID)
	return result, nil
}

func main() {
	// Use a logger for clear output
	logger := log.New(os.Stdout, "[WORKERPOOL]: ", log.LstdFlags)

	// --- 1. Setup ---
	// Create a context for cancellation.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called to free resources

	tenantIDs := []string{"tenant-1", "tenant-2", "tenant-3", "tenant-4", "tenant-5", "tenant-6"}
	numWorkers := runtime.NumCPU() // Use a worker per available CPU core

	// --- 2. Channel Creation ---
	// jobs: Buffered channel to send work to workers.
	// Buffer size equals the number of jobs so the producer isn't blocked. If the buffer size would have been lesser than this,
	// It would result in the produce goroutine which is actually sending the jobs to get blocked, until the receiving end starts
	// Picking up those jobs. But since we know before hand the number of jobs, we can set the buffer size to that value to prevent blocking.
	// This decouples the producers from the workers, since now the producers don't have to be dependent on the workers to receive the jobs,
	// This is the most efficient pattern when the number of jobs are known upfront.
	jobs := make(chan string, len(tenantIDs))

	// results: Buffered channel to receive results from workers.
	// Buffer size also equals the number of jobs. Buffering the results channel ensures that the worker functions doesn't get blocked.
	// For instance, assume the results channel to be unbuffered or of smaller size than the number of jobs. When a worker function tries to send
	// A result through this channel, There is a possibility that the result channel's buffer gets full and now it has to wait for the consumer to
	// Take data from the from the results channel, whereas if the buffer didn't get full, the worker functions wouldn't get blocked, and they could
	// Continue fetching for more jobs and processing them. This maximizes the worker utilization by not keeping them blocked and increases throughput.
	results := make(chan string, len(tenantIDs))

	// errs: Buffered channel for workers to report errors.
	// Buffer size equals the number of workers. In a scenario where the `processData()` function fails and keeps sending errors, the maximum amount of errors
	// That can be receieved are when all the workers receieve an error, and the maximum number of workers is the number of logical processors that the process
	// has access to at the start of the program.
	//
	// If the channel was unbuffered and more than 1 worker tried sending an error down this channel that worker would get blocked until the first error that
	// was sent down this channel is fully processed, decreasing worker function utilization.
	//
	// Setting the buffer size as the number of workers ensures that this program is prepared for the worst case scenario of all workers receiving errors from the
	// process data function and each worker sending those errors down this channel.
	// We ensure that even in this scenario, we don't want the workers to get blocked. Because if the buffer size were any smaller and all workers received errors
	// Then some workers would stay blocked eager to report the errors that they have receieved but they can't until the channel is free to receive errors.

	errs := make(chan error, numWorkers)

	var wg sync.WaitGroup

	logger.Printf("Starting %d workers to process %d jobs...", numWorkers, len(tenantIDs))

	// --- 3. Starting the Worker Pool ---
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(id int, ctx context.Context, jobs <-chan string, results chan<- string, errs chan<- error) {
			defer wg.Done()
			logger.Printf("Worker ID %d has started.", id)

			for {
				select {
				case job, ok := <-jobs:
					// The 'jobs' channel has been closed and drained.
					if !ok {
						logger.Printf("Worker ID %d is shutting down because jobs channel is closed.", id)
						return
					}
					// A job was received. Process it.
					result, err := processTenant(job)
					if err != nil {
						errs <- err
						continue // Skip sending to results and get next job
					}
					results <- result

				case <-ctx.Done():
					// The context was cancelled, signaling a shutdown.
					logger.Printf("Worker ID %d is shutting down due to context cancellation.", id)
					return
				}
			}
		}(w, ctx, jobs, results, errs)
	}

	// --- 4. Distributing Jobs ---
	// This goroutine sends all the jobs to the workers.
	go func() {
		// When there is a single jobs channel being shared by multiple workers, and you distribute jobs down that channel,
		// Go's runtime scheduler would only select one worker and pass the job to that one worker function only.
		// Its a competition and not a broadcast. Its in-deterministic to find out which worker is going to recieve the job.
		// Which is one of the reason to write code in a way that is independent and doesn't change the final outcome when using concurrency.
		//
		// The data passed between channels is passed **ATOMICALLY**.
		for _, tenantID := range tenantIDs {
			jobs <- tenantID
		}
		// Close channel to signal that no more jobs will be sent.
		// The worker function which is ready to accept jobs and process them will receive a boolean value upon closing the jobs channel,
		// that the channel has been closed, and the worker function will exit based on this event.
		close(jobs)
		logger.Println("All jobs have been sent to the jobs channel.")
	}()

	// --- 5. Collecting Results ---
	// This goroutine waits for all workers to finish, then closes the
	// results and error channels to signal that collection is complete.
	go func() {
		// the call to wg.Wait() needs to be in a separate goroutine to prevent Deadlocks.
		// If it was executed inside the main goroutine, then the main goroutine would get blocked waiting for the counter to become 0.
		// But the counter would never change, since neither the results nor the errors channel is getting processed since they are below
		// the this wg.Wait() call. The buffers for the results and the errors channel would get filled completely and waiting for the consume end
		// to receive them.
		//
		// One might think, that putting the wg.Wait() call after the consuming logic from the results and the error channels would solve the issue,
		// Which might seem like it does, but in reality this solution also fails and causes deadlocks.
		// Reason: If you placed the wg.Wait() after the consuming logic of the result and the error channels, you need to think about when would the
		// loop finish in the consuming logic? It would only finish when the results and the error channels are closed using close()
		// But where will you call the close() function now? Inside the workers? You can't since multiple workers in their own goroutines are writing
		// to this channel and an earlier closing of the channel would create errors since another worker might be trying to push their result into this channel
		// So the most omptimal place to close those channels would be when all the workers finish executing, which would technically be when then internal
		// counter of the global weight group hits zero which can only happen after the wg.Wait() call finishes, so now that we have determined that the only
		// place to close the results and the error channels is after the wg.Wait(), we are stuck with another problem. We are placing the calls to close the channels
		// after the consuming logic, which would create a deadlock since the logic which will close the loop is supposed to be executed after the loop ends.
		//
		// To prevent these issues we call wg.Wait() inside another goroutine. This doesn't block the results and errors channels in getting consumed.
		// This entire goroutine executes in parallel with the main goroutine. Its mostly blocked at this wg.Wait() call, until all the workers are finished.
		// After which the results and the errors channel will be safely closed.
		wg.Wait()

		// Why do we close the channels here is also explained in the above gigantic comment. hehe.
		close(results)
		close(errs)
		logger.Println("All workers have finished. Closing results and errors channels.")
	}()

	// --- 6. Processing Outputs ---
	// Read from results and error channels until they are both closed.
	// This loop will block until the goroutine above closes the channels.
	// This consuming logic can be executed on the main thread.
	// This final select loop is designed to drain both channels simultaneously.
	completedJobs := 0
	for completedJobs < len(tenantIDs) {
		select {
		case result, ok := <-results: // Process the result once received.
			if !ok {
				// Set to nil to prevent selecting on a closed channel
				// If the results is not set to nil, then the results channel would get in an always ready state
				// The always ready state of a channel will always be ready for the select statement to select it
				// and it will return the default value of the channel type.
				//
				// This would cause an infinite loop, which will take up 100% of CPU resources, this is called the *Livelock**.
				//
				// By setting a channel in Go as nil, this channel will be **blocked forever** meaning the select statement
				//
				results = nil
				continue
			}
			logger.Printf("RESULT: %s", result)
			completedJobs++
		case err, ok := <-errs: // Process the erro once receieved on this channel.
			if !ok {
				errs = nil // Set to nil to prevent selecting on a closed channel
				continue
			}
			logger.Printf("ERROR: %v", err)
			completedJobs++ // An error still counts as a completed job attempt
		}
		if results == nil && errs == nil {
			break // Both channels are closed and drained.
		}
	}

	logger.Println("Finished processing all jobs.")
}
