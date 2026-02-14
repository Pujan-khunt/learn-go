## Escape Analysis Output
The `-m` commands compiler to ask the compiler where the variables are being put (stack/heap).

```sh
go build -gcflags="-m" main.go

## Output:
# command-line-arguments
./main.go:9:6: can inline square
./main.go:3:6: can inline main
./main.go:5:15: inlining call to square
```

## Benchmarking
```sh
go test -bench=. -benchmem

## Output:
goos: linux
goarch: amd64
pkg: test-01
cpu: 12th Gen Intel(R) Core(TM) i5-1235U
BenchmarkSquare-12    	1000000000	         0.2142 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	test-01	0.235s
```

The `0 allocs/op` indicates there were total of 0 _allocations_ per _operation_
Hence it verifies the escape analysis that no variable escapes to the heap.
