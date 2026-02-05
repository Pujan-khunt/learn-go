package main

func main() {
	ch := make(chan int)

	// go func() {
	// 	for val := range ch {
	// 		println(val)
	// 	}
	// }()

	for i := range 10 {
		ch <- i + 1
	}
}

