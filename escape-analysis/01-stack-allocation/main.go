package main

func main() {
	n := 4
	sqr := square(n)
	println(sqr)
}

func square(n int) int {
	val := n * n
	return val
}
