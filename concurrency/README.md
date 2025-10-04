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
t = 4.0s: Sending message every 2s
