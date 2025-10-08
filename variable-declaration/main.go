package main

import "fmt"

func main() {
	// Constant variables can be typed or untyped.
	// Constant Declaration.
	// Immutable Value Variable = Constant.
	const val = 123                   // untyped -> can be used in any context which requires an integer like int32 int64 etc.
	const name string = "Pujan Khunt" // typed -> can be only used in context requiring string type.

	// This gives an error since constant values need to be known at compile time, rather than the runtime.
	// const Greeting string = fmt.Sprintf("Hello %s", name)

	// Grouped Constants
	const (
		StatusOk       = 200
		StatusNotFound = 404
		Version        = "1.0.0"
	)

	// iota identifier is a predefined constant which simplifies declaring of sequential constants.
	// The value of iota is reset to 0 in every const statement/block.

	// Example
	const a = iota // value of a = 0

	const b = iota // value of iota resets to zero and is assigned to be. Therefore b = 0.

	const (
		a0 = iota      // iota resets to zero, a0 = 0
		a1 = iota      // iota doesn't reset since its in the same context of const block. Therefore iota is incremented by 1. a1 = 1
		a2 = iota      // a2 = 2
		a3 = iota * 10 // a3 = iota * 10 = 3 * 10 = 30
	)

	// Normal/Dynamic variable creation using var keyword
	// Unlike const, var is for runtime values.

	var variable int // syntax without type inference

	// variables created using var have default values when created without values.
	// 0 for numeric types (int, float64)
	// false for booleans
	// ""(empty string) from strings
	// nil for reference types like slices, maps, channels also for pointers(i think)

	// grouped variable declaration using var keyword
	var (
		var1 string
		var2 map[int]string
		var3 <-chan bool
	)

	// syntax with type inference
	version := 1.22

	// shorthand for above declaration
	version2 := 1.22 // can't use the same variable name because of conflict

	// Printing each variable to avoid getting unused variable errors.
	fmt.Println(variable, var1, var2, var3, version, version2)
}
