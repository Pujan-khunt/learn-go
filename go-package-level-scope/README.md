Since I am new to Golang, I was checking out the package level scope that Golang has.

So basically what I learnt is that if you create any one of the following in a separate file
which is in the same package as the other file, then that other file can easily use these things
without importing anything. 
- function
- variable
- struct

> NOTE: Also the convention is to group files having the same package in the same directory.

Hence the files in the same directory don't need to import anything if they want to use something from the other files.

## How to run these files?
If you just run

```bash
go run main.go
```

This will give you an error, since it doesn't know about the functions, variables and structs which are in other files but in the same package.
Hence you need to include each file

```bash
go run main.go test.go
```

or instead you can just do.

```bash
go run *.go
```

or even first build and then run

```bash
go build
```

If you need to use any `function`, `variable` or `struct` from any of the files but from a file belonging to a different package, 
then you need to import that package and also must ensure that the thing you are importing is **capitalized**.

- Capitalized: Exported. Other packages (directories) can import it and use it.
- Lowercase: Unexported (Private). Only visible to the files inside the same package (directory).

## How to import attributes (functions, variables and structs) from other package

Here is where the `go.mod` file shines.

To use code from a different package (directory), you must import it. This gives the project a `namespace`.

You define the module using the following command:
```bash
go mod init <module-name>
```

The name of the current module is `test-module`.

The above command creates a `go.mod` file with its first line as.
```gomod
module <module-name>

go <go-version>
```

Whenever you write an import statement in a go file, you are importing a go package.
