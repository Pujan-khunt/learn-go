## Example for using an external package not in the standard library

Search for external go packages here: [Go package registry](pkg.go.dev)

> We are importing the rsc.io/quote/ package version 4.

Link to the [package](pkg.go.dev/rsc.io/quote/v4#Go)

The `documentation` section contains an `index` section.
Under the `index` section, you can see all the function that the package provides you to use.

We are going to use the `Go` function.

Import the external package using the `import` statement.

After importing you need to install it as a dependency. To do that, run the following command:
```bash
# Installs the rsc.io/quote dependency
go mod tidy
```
