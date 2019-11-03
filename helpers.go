// rubberduck/helpers.go

package main

import (
	"os"
	"time"
)

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

// Exists takes a filepath string and checks if it exists, returning the appropriate boolean
func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func appendNewline(s string, n int) string {
	for i := 0; i < n; i++ {
		s += "\n"
	}
	return s
}
