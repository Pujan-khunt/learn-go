package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	str := "Hello"
	integer := 12345
	fmt.Println(reverse.String(str))
	fmt.Println(reverse.Int(integer))
}
