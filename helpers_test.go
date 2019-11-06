// rubberduck/helpers_test.go

package main

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	file := "helpers_test.go"
	if !Exists(file) {
		t.Error("Expected helpers_test.go to exist (since that's, y'know, me, and Descartes seems like a smart dude).")
	}
}

func ExampleExists() {
	file := "helpers_test.go"
	fmt.Println(Exists(file))
	// Output:
	// true
}

func BenchmarkExists(b *testing.B) {
	file := "helpers_test.go"
	for i := 0; i < b.N; i++ {
		Exists(file)
	}
}
