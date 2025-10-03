# `defer` keyword in Golang

> `defer` in golang is used to ensure that a function call is executed later 
> in a program's execution, usually for the purpose of cleanup.
> `defer` is usually similar to the `finally` keyword in Java.

Basically defer schedules a function call to be run at the end of a function's execution.
```go
func functionOrMethodCall(a int) {
    fmt.Println(a)
}

func main() {
    a := 1
    fmt.Println("Start.")
    defer functionOrMethodCall(a) // When stack pointer reaches here, it doesn't execute this function but binds the value of a to 1
    a = 2 // Value of a is updated here, but doesn't change the argument being passed in the above function call.
    fmt.Println("End.")
}
```

## Two Critical Points.
1. **Argument Evaluation is Immediate**: The arguments passed to the deferred function are evaluated in the order
that the stack pointer moves, only the function call is delayed, the value that is passed as arguments is decided
on when the stack pointer is at the defer keyword line.

2. **Execution Order is LIFO**: If a function has multiple `defer` statements, then their execution would happen,
in Last-In-First-Out fashion. 

> Its simple to imagine that first the stack-pointer moves from top to bottom, then 
for the defer statements, it moves from bottom to top.


## Usecases
1. The primary use case is to **guarantee resource cleanup**. It ensures that the resource created will be 
cleaned up at the end of the execution, whether its an early return, panic, or even a normal return.

```go
func processFile(fileName string) error {
    f, err := os.Open(fileName)
    if err != nil {
        return err
    }
    defer f.Close() // Guarenteed to run before the function returns or panics

    // process the file
    //...
    //...
    //...

    return nil
}
```

2. Unlocking mutexes, ensuring no deadlocks.
```go
import "sync"

var mutex sync.Mutex
var data map[string]string

func writeToMap(key, value string) {
    mutex.Lock()
    defer mutex.Unlock()

    data[key] = value

    // even after a panic occurs here, mutex will still be unlocked.
}
```

3. Closing DB connections.
```go
func getAlbums(db *sql.DB) error {
    rows, err := db.Query("SELECT * FROM album")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        // process and convert rows of data into a strongly typed object.
    }

    return rows.Err()
}
```

## How Go Handles it under the hood?
> Go manages defer statments using a stack like data-structure associated with each goroutine.

When a function is called, a dedicated stack space is created for the function call.
The stack pointer starts from the top of the body of the function and starts executing each line, line by line.

When the stack pointer reaches a defer statement, it does 2 things:
1. Evaluates the arguments to the deferred function call.
2. Push the pointer to the function inside the defer stack.

When the main outer function is finished executing, either via body exhaustion or early return, normal return or even a panic,
Now its time to check the defer stack and start popping off each function from the stack in LIFO order and start executing them.

Each stack space for a function call would have their own defer stack.
