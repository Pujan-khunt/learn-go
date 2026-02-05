package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	// Get file handler with creating if it doesn't exist.
	f, err := os.OpenFile("data.bin", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Ensure file has a set size (cannot map an empty file)
	size := int64(1024) // 1 KB
	if err := f.Truncate(size); err != nil {
		log.Fatal(err)
	}

	data, err := syscall.Mmap(int(f.Fd()), 0, int(size), syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal(err)
	}
	defer syscall.Munmap(data)

	msg := "Hello, memory mapped world!"
	copy(data, msg)

	data[0] = 'h'

	fmt.Println("Data written to memory map. Check the memory mapped file contents.")
}
