package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

/*
Imagine you need to process a 10 GB file, but your RAM is only 8GB.

This means that you can't load the entire file into your RAM, because of a hardware limit.
The solution is to use memory mappiing where the kernel provides you the abstraction to
manipulate any part of that 10 GB file but doesn't actually load into the memory, rather it just
updates its page table to allow your process the access to using the virtual memory address of the
file located in your disk of size larger than 10 GB.

Now when your go code is running as CPU instructions it will try to use that virtual memory address
assigned by your kernel, and the CPU will halt this process and will give a PAGE FAULT INTERRUPT, since
it will using the MMU (Memory Management Unit) that the virtual address isn't being backed by any RAM,
i.e. There is no physical address on the RAM corresponding to the virtual memory address that is provided
by the process.

Now our go process would stop and the kernel would take over, the kernel would look into its page table and
identify that the virtual memory address that our process is trying to access is located in a specific page in
the memory mapped file, it will load that specific and single page of 4KB into the RAM, now it would resume our
process and also update its page table to point that virtual memory address to the RAM where it loaded that 4KB page.

NOTE: Your file explorer or the `ls -lh` command might say that the file size is 1 GB which is true and false.
It says the size is 1GB because that is the apparent size of the file. Which means that you can pull 1 Giga bytes out of these file,
even though they don't exist, when you will try to read any middle portion of this file, the File System sees it as `unallocated` and
will generate the zeroes for you on the fly.

If you try to write in the middle later, the File System will allocate a real storage block for that specific spot at that moment.

If you want to see how much space is it actually taking up, then run the following command.

```bash
du -h large_sparse.data
```

This command should show you '4.0K' which means the size of the single page that it created when modifying the last index which is 4 KiloBytes.

Comparision:
read() -> 10GB in RAM
mmap() -> 4KB in RAM

1 GB (gigabytes) = 1024 MB (megabytes) = 1024 * 1024 KB (kilobytes) = 1024^3 B (bytes) = 8 * 1024^3 b (bits)
*/
func main() {
	const OneGB = 1 << 30
	filename := "large_sparse.data"

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Error creating file: %s\n%v\n", filename, err)
	}
	defer f.Close()

	if err := f.Truncate(OneGB); err != nil {
		log.Fatalf("Error truncating file to 1 GB: %s\n%+v\n", filename, err)
	}

	data, err := syscall.Mmap(int(f.Fd()), 0, OneGB, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatalf("Error memory mapping file: %s\n%+v\n", filename, err)
	}
	defer syscall.Munmap(data)

	// Since we have random access, we can jump to the last byte of the 1GB file.
	// The kernel would only load the page containing this specific byte into the RAM.
	// It does NOT load the empty gigabyte before it.
	index := OneGB - 1
	data[index] = 'Z'

	fmt.Printf("File size is logically: %d bytes\n", len(data))
	fmt.Printf("Written 'Z' at index: %d\n", index)
	fmt.Println("Check the file size on disk! It might look huge but use little space.")
}
