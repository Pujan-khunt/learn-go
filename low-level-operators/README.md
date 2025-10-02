## Memory Layout:
When a go program runs, it uses two main regions of memory to store data:

### 1. Stack Memory
Highly organized and efficient region of memory. When a function is called,
it gets its own block of memory called **stack frame** which is reserved on top of the stack frame.

All of the local variables of the function are stored in this **stack frame**. When the function
exits, the frame is popped off the stack and the memory is automatically reclaimed.

The stack is fast because it only involves a single **stack pointer**.

### 2. Heap Memory
This is a much larger section, far less organized region of memory data,
which is required when you need to store data which needs to outlive a 
single function call.

Managing heap is more complex and is managed by the GC of Golang.

## Golang consists of 2 low level operators

### 1. & Address Of Operator
The & operator receives a variable and it returns a unique address,
indicating the location where the variable is stored in memory (RAM).
In x86 Assembly, this is essentially done by the `LEA` (Load Effective Address) Instruction.
It doesn't read the data, just gets the location.

> LEA calculates the memory address of the source operand and places its value into the destination register.

### 2. * Dereference Operator
A pointer is a **special** kind of a variable which holds a **memory address**.
Now you use the * operator with a pointer, which basically is the inverse of the & operator.
* takes the memory address stored in the pointer, goes at the memory address and brings back the value.
In x86 Assembly, this is essentially done by the `MOV` (MOVE) Instruction.
Example:

Gets the value at register `b` which is a memory address, goes to the memory address 
and receives the value and moves it into the register `a`.
```nasm
MOV a, (b)
```


## Generating Assembly Code from Golang Source Code
```bash
go tool compile -S main.go
```
