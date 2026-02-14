```sh
go build -gcflags="-m" main.go

## Output:
# command-line-arguments
stack-allocation-due-to-inlining/main.go:8:6: can inline square
stack-allocation-due-to-inlining/main.go:3:6: can inline main
stack-allocation-due-to-inlining/main.go:5:8: inlining call to square
stack-allocation-due-to-inlining/main.go:8:13: n does not escape
```

Here `n` does not escape to the heap and stays on stack.
Reason: Since there is no use of n after the function call to `square()`
the function does not need to persist the value of `n` between function stack frames.
Hence it stays on the stack.
