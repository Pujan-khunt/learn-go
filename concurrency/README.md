# Concurrency in Golang

**Concurrency**: programs which are executed in out-of-order or in a partial order and don't affect the final outcome and are independently executed are said to achieve Concurrency.

**Parallelism**: Its about doing a lot of things at once at the exact same moment in time. Eg. splitting a data processing task over to multiple CPU cores.

In the context of Go, **Goroutines** help us achieve both **Concurrency** and **Parallelism**.
Goroutines help us to write code which is independent of code from other Goroutines, they are not inherently parallel.
By default, Go only uses a single OS thread, regardless of the number of Goroutines.
However, with the help of the **Go Runtime Scheduler** and `GOMAXPROCS` set to the number of logical processors, we can achieve true parallel execution of Goroutines.


## Advantages and Disadvantages of Concurrency

### Advantages:
1. **Resource Efficiency**: Goroutines are lightweight by nature. Unlike traditional OS threads, which require a significant memory allocation for stack space(often measured in megabytes), Goroutines start with a much smaller stack space(often measured in kilobytes) which can adapt in size as required. Due to this small initial size, you can create a large number of Goroutines(even in millions), without exhausting memory, thereby boosting resource efficiency.

2. **Synchronization Primitives**: Go's concurrency models avoids manual lock based synchronization and rather provides high level constructs which are less error-prone. This avoids many common problems like deadlocks, livelocks, and race conditions when working with lock based synchronization.

3. **Go's Standard Library**: Go SL(Standard Library) contains many packages to assist in concurrent programming. For example, the `sync` package provides synchronization primitives like `WaitGroup` and `Once`. The `sync/atomic` package provides many low-level atomic memory operations, allowing lock-free concurrent programming.

### Disadvantages:
1. **Concurrency Is Not Parallelism**: While Goroutines can achieve concurrency, true parallelism is dependent on the Go runtime's ability to distribute Goroutines into multiple/individual CPU cores, which isn't always guarenteed.

2. **Shared Data and Data Races**: Shared state mutation is still a possibility. This can lead to data races when multiple Goroutines trying to access the shared data without proper synchronization.

3. **Debugging and Profiling**: Debugging and profiling Go applications can be complex, due to the fact that the behaviour of Goroutines is indeterministic. Go provides the `pprof` package and a race detector, but understanding these tools and using them effectively requires much understanding and time.


## Goroutines
> A Goroutine is a lightweight thread of execution. Goroutines are functions/methods which concurrently run with other Goroutines.

```go
package main

import (
    "fmt"
    "time"
)

func printMessage() {
    fmt.Println("Hello from a Goroutine.")
}

func main() {
    go printMessage()
    fmt.Println("Hello from the main function.")
    // Wait for the Goroutine to finish.
    time.Sleep(time.Second)
}
```
We are running the `printMessage()` function inside a separate Goroutine. Both `printMessage()` and `main()` are being run concurrently.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	count("sheep")
	count("fish")
}

func count(animal string) {
	for i := 1; true; i++ {
		fmt.Printf("%d %s", i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}
```

For this example, the `count("fish")` function will never be called since the `count("sheep")` will never terminate due to the infinite loop.


```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")
}

func count(animal string) {
	for i := 1; true; i++ {
		fmt.Printf("%d %s", i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}
```

The `go` keyword here creates a new Goroutine which and the sole task of that Goroutine is to execute that function call.
We can create a Goroutine since both tasks (both function calls) are independent of each other and don't change the final outcome.

Here the flow would go as, the stack pointer would begin executing the main function and would see the go keyword and will create a new
Goroutine which will run concurrently as the main Goroutine would continue running the rest of the main function concurrently, i.e. executing
the `count("fish")` function call.

> A very interesting thing happens in this example.
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	go count("fish")
}

func count(animal string) {
	for i := 1; true; i++ {
		fmt.Printf("%d %s", i, animal)
		time.Sleep(time.Millisecond * 500)
	}
}
```

Now that both function calls are being handled by separate Goroutines, both should run concurrently, right? Kind of.
If you execute the above example code, you will find that nothing is printed to the stdout.

Reason: Main Goroutine finishes executing once it creates 2 new Goroutines for executing both function calls.
The job of the main Goroutine is to execute the `main()` function and it does that pretty quickly here and finishes.

In Go, once the main Goroutine finishes, the program/process terminates no matter what the other Goroutines are doing. 
In the other examples the main Goroutine was always busy doing something and hence the program wasn't terminating instantly.

To actually make the above example code work, we need to make the main function busy, that can be done by something like this:
```go
time.Sleep(time.Second * 100)
```

This will keep the main function busy for 100 seconds and the loops will run for approximately 200 iterations.

Another approach could be to ask for user input, since the go process will wait for the user to input something, as long as the user doesn't do that, 
the main function is busy waiting.
```go
fmt.Scanln()
```

Both of above techniques can only be used for testing purposes, since they are indeterministic and unusable for actual production code.
The most reccommended way is to use a `WaitGroup`

## Channels
> Goroutines execute tasks concurrently. Channels provide a way to synchronize and control these tasks.
> Channels are a way through which you can send and receive values using the channel operator `<-`.

```go
package main

import (
    "fmt"
    "time"
)

func printMessage(message chan string) {
    time.Sleep(time.Second * 2)
    message <- "Hello from Goroutine"
}

func main() {
    message := make(chan string)
    go printMessage(message)
    fmt.Println("Hello from main function.")
    fmt.Println(<-message)
}
```

In this example, the `printMessage()` function waits for 2 seconds, then sends the message on the channel.
The main function is running concurrently while this is happening and prints the `Hello from main function.` message 
and then receives the message from the channel.


## Wait Groups
A Wait Group waits for Goroutines to finish. It's a struct type and maintains a counter, which represents the active Goroutines actively executing.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Creating a global wait group.
var wg sync.WaitGroup

func worker(id int) {
    // Mark the Goroutine as completed which will update the internal counter, and as soon as the 
    // counter of the global wg (wait group) reaches zero, main goroutine is terminated.
    defer wg.Done() 
    fmt.Printf("Worker %d starting\n", id)
     // pausing the Goroutine to mimic execution of worker.
    time.Sleep(time.Second)
    fmt.Printf("Workder %d is done\n.", id)
}

func main() {
    for i := 1; i <= 5; i++ {
        wg.Add(1) // Increments the counter to indicate that a Goroutine has been created.
        go worker(i)
    }
    wg.Wait() // Waits until the internal counter becomes zero.
}
```

In this above code example, we add 1 to the `wg` counter every time we create a new Goroutine.
Then call `wg.Done()` once the Goroutine function is finished executing, which decrements the counter of `wg`.
`wg.Wait()` will block the main goroutine until the counter reacher zero.

Wait groups are simple counters, which are simple but not massively beneficial, until we introduce **channels**.


## Channels
A channel is a way for Goroutines to communicate with each other. Till now, we are only printing from the worker function that is executed by the Goroutine,
but what if we wanted to send data back to the main goroutine which called it. Then we use channels.

A channel is like a pipe through which you can send a message or receive a message.

Accept channel as an argument in the function which the goroutine will execute.

Channels have a type as well. For eg. a `string channel` can only send or receive data of type `string`

> **Fun Fact**: the data type of a channel can be another channel too.

> **NOTE**: sending and receiving to and from a channel is blocking. Meaning the goroutine will be blocked until the other party does thier job.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // Create a channel of type string
	go count("sheep", c)   // Create a Goroutine passing the function arguments along with a channel of type string

	// If the msg: <- c is still trying to receive a message but the goroutine has already finished then it will result in a deadlock. Try converting the below for loop into an infinite one, see it in action (only works without a `close()`).
	for i := 1; i <= 10; i++ {
		// Create a variable of type string to receive string output from the above created goroutine's output
		// When receiving from a channel, you can create 2 variables, one which will store the data from the channel which must of the same type as the channel
		// Second is a boolean indicating if the channel is open or not.
		msg, open := <-c

		if open {
			fmt.Println(msg)
		} else {
			fmt.Println("Channel is closed: ", i)
		}

	}
	// There exists a much cleaner syntax of this version
	// Here instead of checking for the open and receiving the channel data you can loop over the range of a channel, which will essentially do the same thing.
	// for msg := range c {
	// 	fmt.Println(msg)
	// }
}

func count(animal string, c chan string) {
	for i := 1; i <= 5; i++ {
		// fmt.Printf("%d %s", i, animal)
		c <- fmt.Sprintf("%d %s", i, animal) // Sends the output from the format.Sprintf() command into the channel.
		time.Sleep(time.Millisecond * 500)
	}

	// Never close the channel as a receiver. Here its ok since we are at the sender's end.
	// WHY? If the sender tries to send the data after receiver closes the channel, the program will panic.
	close(c) // Close the channel so the receiving end doesn't get deadlocked when still trying to receive even after the goroutine has ended.
}
```

A deadlock situation in the below code.

```go
package main

import "fmt"

func main() {
    c := make(chan string)
    c <- "Hello World!"

    msg := <- c
    fmt.Println(msg)
}
```

One might think, that the above code, creates a channel, sends the "Hello World!" string 
and receives that same string in the same Goroutine and prints to the stdout.
But you would notice after running this code, you are getting a deadlock.

**WHY?**
The same reason as mentioned in the above NOTE. Sending and receiving to and from a channel is blocking.
Hence the main goroutine will be blocked at the line `c <- "Hello World!` trying to wait for another goroutine
to consume that string that the main goroutine just sent this channel.

## Select Keyword

### Why is Select keyword needed?
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    // Goroutine which is sends a message to the c1 channel every 500ms
    go func() {
        for {
            time.Sleep(time.Millisecond * 500)
            c1 <- "Sending message every 500ms"
        }
    }

    // Goroutine which is sends a message to the c2 channel every 2s
    go func() {
        for {
            time.Sleep(time.Second* 2)
            c1 <- "Sending message every 2s"
        }
    }

    for {
        fmt.Println(<- c1)
        fmt.Println(<- c2)
    }
}
```

The problem that will be faced in this above example code is that:
t = 0.5s: Sending message every 500ms
t = 2.0s: Sending message every 2s
t = 2.5s: Sending message every 500ms
t = 4.0s: Sending message every 2s
t = 4.5s: Sending message every 500ms
t = 6.0s: Sending message every 2s

Even though the first goroutine is capable of sending messages every 500ms its getting blocked because of the slower goroutine.

To solve this issue, we use `select` statement.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine which is sends a message to the c1 channel every 500ms
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Sending message every 500ms"
		}
	}()

	// Goroutine which is sends a message to the c2 channel every 2s
	go func() {
		for {
			time.Sleep(time.Second * 2)
			c1 <- "Sending message every 2s"
		}
	}()

	// Infinite loop
	for {
		// Whenever there is a channel ready to transmit data, execute the code below that case statement.
		// Not the fast Goroutine will not be blocked by the slow Goroutine.
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
```

Now the output will look like:
t = 0.5s: Sending message every 500ms
t = 1.0s: Sending message every 500ms
t = 1.5s: Sending message every 500ms
t = 2.0s: Sending message every 500ms
t = 2.0s: Sending message every 2s
t = 2.5s: Sending message every 500ms
t = 3.0s: Sending message every 500ms
t = 3.5s: Sending message every 500ms
t = 4.0s: Sending message every 500ms
t = 2.0s: Sending message every 2s


## Worker Pattern Using Oneway Channels

### Oneway channels: you can define a one way channel in 2 ways:
1. **Only Send**: Channel which can only send data, but cannot receive data.

```go
func fn(c <-chan string) {}

```
2. **Only Receive**: Channel which can only receive data, but cannot send data.

```go
func fn(c chan<- string)
```

Worker pattern uses these oneway channels.

```go
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
```

The entire working and logic is explained in the above code in the form of comments.

There can also be another performance boost in the case of slow logic processors.
Imagine this scenario where the processing that happens after receiving the result from the results channel, is very slow.

You have 2 options to solve this issue:
1. Create separate a goroutine for each processing task
2. Put the entire consuming logic inside a goroutine itself.
