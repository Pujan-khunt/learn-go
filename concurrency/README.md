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
