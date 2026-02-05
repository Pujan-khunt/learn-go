package main

import "time"

func main() {
	ch := make(chan int, 10)

	for i := range 10 {
		ch <- i
	}

	go func() {
		for val := range ch {
			println(val)
		}
	}()

	time.Sleep(1 * time.Second)
}
