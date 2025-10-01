Go code is grouped into `packages` and `packages` are grouped into `modules`. 
Go module defines the dependencies needed to run the code, including the go version.

Add comment in the format of `<function-name> <comment>` to add it in the docs for a specific function.
Same for packages but the format is `Package <package-name> <comment>`. Hover to view the docs.

A function starting with a capital letter is exported, meaning it can be called by a file not in the same package.
This is known as a `exported name`.

Add error handling using the `errors` standard library package from go.
```go
errors.New("error message")
```

Implement random greeting message from a slice
use the `math/rand` package
