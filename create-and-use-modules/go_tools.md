# Building a Go project
```bash
go build -o <executable-name>
```
1. Creates an executable for projects with `package main` declarations
2. Compiles the code and stores in the build cache for `library` packages.

go projects with a `package main` declaration produce an `executable` based on the following env variables used by go:
1. GOARCH
2. GOOS

To view the values of these env variables use the following command:
```bash
go env # To view all go env variables

go env GOARCH # To view a specific env variable

go env -w GOBIN $(go list -f '{{.Target}}') # To permanently set a env variable
```

# Install a Go project
```bash
go install
```

go install command is very similar to the go build command.

the install command first does everything that the go build command does.
Then the binary executable generated (if generated) is placed into the `$GOPATH/bin` or `$GOBIN`

If the above mentioned directory is in your system's path then after running go install, you can 
run the executables anywhere from your terminal.

## How to add $GOBIN to path.
```bash
printf "export PATH=$PATH:/path/to/gobin/directory\n" >> ~/.bash_profile
```

# Usage of go install command
1. Installing projects written in go from GitHub
2. installing your custom projects written in go for testing purposes


# Go mod command
> Mainly for **dependency management** using Go Modules. Interacts with the **go.mod** file.

1. go mod init [module-path]
Creates a new module in the current directory, by creating a go.mod file

2. go mod tidy
Synchronizes the dependencies by removing unnecessary ones from the go.mod file 
and adding the necessary ones by looking in the source code.

# Go work command
Enables workspace mode in a go project. Makes working with multiple modules simpler.

- Allows to develop and test code in multiple modules at once without needing to edit the go.mod file
(i.e. alternative for using replace directive)

- It uses a go.work file for listing all the modules which are part of the workspace. The go toolchain
will now use the local versions if available in the workspace rather then getting from a remote registry.
