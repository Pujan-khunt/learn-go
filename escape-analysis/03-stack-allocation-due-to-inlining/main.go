package main

func main() {
	n := 4
	square(&n)
}

func square(n *int) {
	*n *= *n
}

