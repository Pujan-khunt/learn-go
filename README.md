# First go project

Testing out how to compile, run, format and configure lsp with Go.
1. Creating a go module
```bash
# Creates a go module with the go.mod file
# The name of the module is "example.com/hello"
# Generally it is the name of the repository
go mod init example.com/myproject
```

2. Running go project
```bash
# Runs the package "main"
# It doesn't store the binary, it just runs it.
go run .
```

3. To build/compile the go project
```bash
# Builds the main package.
# Name will be same as your folder name
go build .
```

4. Run tests (all)
```bash
# Runs all go test files
go test ./...
```


# go.mod file

> When importing dependencies from other modules, you manage those dependencies through your package's own module.
> Your module is defined in the **mod.go** file. It tracks the modules that provides those packages. **mod.go** stays
> with your source code in your vcs.



# Go tools
1. Go binary
pretty much a all in one tool, installer, compiler, tester etc.

2. gopls
LSP for Golang

3. gofumpt
Strict go code formatter

4. delve
Go debugger
