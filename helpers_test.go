// rubberduck/helpers_test.go

package main

import (
	"fmt"
	"testing"
)

func TestExists(t *testing.T) {
	file := "helpers_test.go"
	if !Exists(file) {
		t.Error("Expected Exists(helpers_test.go) to exist.")
	}
}

func TestMin(t *testing.T) {
	x := 1
	y := 10
	m := min(x, y)
	if m != x {
		t.Error("Expected min(1, 10) to return 1. (Returned ", m, ")")
	}
}

func TestAppendNewline(t *testing.T) {
	s := "Hello there."
	n := 1
	x := appendNewline(s, n)
	if x != s+"\n" {
		t.Error("Expected last character of x to be a newline after appendNewline was called")
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

func BenchmarkMin(b *testing.B) {
	x := 11
	for y := 0; y < b.N; y++ {
		min(x, y)
	}
}

func BenchmarkAppendNewline(b *testing.B) {
	s := "Hello there."
	n := 2
	for i := 0; i < b.N; i++ {
		appendNewline(s, n)
	}
}
