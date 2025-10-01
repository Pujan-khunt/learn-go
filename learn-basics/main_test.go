package main

import "testing"

func TestMain(t *testing.T) {
	if true {
		t.Errorf("This test failed")
	}
}
