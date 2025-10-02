package main

import "fmt"

/*
* The `go:inline` directive disables smart optimizations to get applies by the Go compiler
* This ensures that the function looks similar to the source code in Assembly.
 */

//go:noinline
func passByValue(x int) {
	x = 20
}

//go:noinline
func passByPointer(ptr *int) {
	*ptr = 30
}

func main() {
	a := 10
	passByValue(a)
	fmt.Println(a)

	b := 20
	passByPointer(&b)
	fmt.Println(b)
}
