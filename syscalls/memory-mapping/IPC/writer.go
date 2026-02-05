package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	f, _ := os.OpenFile("shared_mem.bin", os.O_RDWR|os.O_CREATE, 0644)
	f.Truncate(1024)
	defer f.Close()

	data, _ := syscall.Mmap(int(f.Fd()), 0, 1024, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	defer syscall.Munmap(data)

	fmt.Println("Writer started. Updating counter in shared memory...")

	for i := range 100 {
		binary.LittleEndian.PutUint64(data[:8], uint64(i))
		fmt.Printf("Wrote: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
