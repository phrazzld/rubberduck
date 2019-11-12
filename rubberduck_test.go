// rubberduck/rubberduck_test.go

package main

import (
	"testing"
	"time"
)

func TestInitDatetime(t *testing.T) {
	startTime := time.Date(2011, time.December, 1, 2, 3, 4, 0, time.UTC)
	expectedDate := "2011 December 1"
	expectedTime := "02:03:04"
	actualDate, actualTime := initDatetime(startTime)
	if actualDate != expectedDate {
		t.Errorf("(expected) %s (actual) %s", expectedDate, actualDate)
	}
	if actualTime != expectedTime {
		t.Errorf("(expected) %s (actual) %s", expectedTime, actualTime)
	}
}
