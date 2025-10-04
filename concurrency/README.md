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



