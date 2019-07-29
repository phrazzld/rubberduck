// Rubberduck tests, examples, and benchmarks
package main

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	file := "stub"
	if !Exists(file) {
		t.Error("got something wanted something else")
	}
}

func ExampleExists() {
	file := "stub"
	fmt.Println(Exists(file))
	// Output:
	// true
}

func BenchmarkExists(b *testing.B) {
	file := "stub"
	for i := 0; i < b.N; i++ {
		Exists(file)
	}
}
