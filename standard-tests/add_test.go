package add

import "testing"

func TestAdd(t *testing.T) {
	inputA := 3
	inputB := 2
	expected := 5

	result := Add(inputA, inputB)

	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", inputA, inputB, result, expected)
	}
}

// If the test includes multiple testcases, writing if statements for all can be messy.
// Hence go uses table driven tests. This is the standard.
// You define a slice of testcases and iterate over them.
func TestAdd_TableDriven(t *testing.T) {
	// Define the list of testcases
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"Positive numbers", 2, 3, 5},
		{"Negative numbers", -1, -2, -3},
		{"Mixed numbers", -5, 5, 0},
		{"Zeros", 0, 0, 0},
	}

	// Loop over all testcases and execute them.
	for _, tc := range tests {
		// t.Run creates a sub-test. If any one sub-test fails, others still run.
		t.Run(tc.name, func(t *testing.T) {
			got := Add(tc.a, tc.b)

			if got != tc.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
