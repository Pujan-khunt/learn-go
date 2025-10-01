For production cases, you would publish the go module `greetings` for usage, but here we have to import the local go module.

To do that, edit the go tools to find the local path to the go module `greetings`
```bash
go mod edit -replace example.com/greetings=../greetings/
```

Another solution exists where you can use `workspaces` in go. Create a `go.work` file at the root level (above both modules)
```bash
<root-folder>
├── go.work
├── greetings
└── hello
```

Create this `go.work` file by running the following command and specifying all the modules
```bash
go work init ./greetings ./hello
```
