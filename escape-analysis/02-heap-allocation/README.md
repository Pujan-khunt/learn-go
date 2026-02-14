## Escape Analysis

```sh
go build -gcflags="-m" main.go

## Output:
# command-line-arguments
heap-allocation/main.go:9:6: can inline square
heap-allocation/main.go:3:6: can inline main
heap-allocation/main.go:5:15: inlining call to square
heap-allocation/main.go:10:2: moved to heap: val
```

Here we can see that since we returned a pointer from `square()` the pointer to `val` was moved to the heap. 
The `val` pointer needs to be persisted between function calls and stack frames
since its going to be used in the next function call to `println()`.

## Benchmarking

```sh
go test -bench=. -benchmem

## Output:
goos: linux
goarch: amd64
pkg: test-02
cpu: 12th Gen Intel(R) Core(TM) i5-1235U
BenchmarkSquare-12    	134965215	         8.979 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	test-02	2.117s
```

The `1 allocs/op` indicates a total of 1 memory allocation per operation. The size of that allocation is `8 Bytes` per operation (which is the size of `int` in 64-bit systems).
This verifies the escape analysis that `val` escapes to heap.
