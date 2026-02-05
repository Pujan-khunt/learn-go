package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := stat.Size()

	data, err := syscall.Mmap(int(f.Fd()), 0, int(size), syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Munmap(data)

	if size > 0 {
		fmt.Printf("First byte: %c\n", data[0])
		fmt.Printf("Content: %s\n", string(data))
	}
}
