the folder `4.multi-module-workspaces` is the `workspace` folder and the `hello` is a module inside the workspace.

The `go.work` file represents information about the workspace.
It also mentions the modules inside the workspace.
`hello` is one of them.

The `go.work` file similar to the `go.mod` file contains the go directive which tells Go which version of Go should the file be interpreted with.

Now, the following command can run the module

Adding the local module (cloned from googlesource) to the workspace using:
```bash
go work use ./example/hello
```

The workspace now includes the ./hello module and the ./example/hello module which provides the golang.org/x/example/hello/reverse package.
This will allow to write new code in our copy(clone) of the reverse package.

```bash
# Finds the main package inside the hello module and runs the main function.
go run ./hello # From the workspace directory
```
