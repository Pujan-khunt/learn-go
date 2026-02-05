package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	f, _ := os.Open("shared_mem.bin")
	defer f.Close()

	// Map the SAME file
	data, _ := syscall.Mmap(int(f.Fd()), 0, 1024, syscall.PROT_READ, syscall.MAP_SHARED)
	defer syscall.Munmap(data)

	fmt.Println("Reader monitoring shared memory...")

	for {
		// Read the first 8 bytes directly from memory
		val := binary.LittleEndian.Uint64(data[:8])
		fmt.Printf("\rCurrent Value in Memory: %d", val)
		time.Sleep(500 * time.Millisecond)
	}
}
