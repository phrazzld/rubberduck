// rubberduck/search.go

package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func search(notesPath string, searchTerm []string) (hits []string, err error) {
	var b strings.Builder
	// Turn `rubberduck search <term>` into:
	// egrep <searchTerm> -R -w <notesPath> --ignore-case -C 1
	cmd := exec.Command(
		"egrep",
		strings.Join(searchTerm, " "),
		"-R",
		"-w",
		notesPath,
		"--ignore-case",
		"-C",
		"1",
	)
	cmd.Stdin = os.Stdin
	// Write output to strings.Builder
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return hits, err
	}
	// Break each hit into a separate item in our returned slice of strings
	hits = strings.Split(b.String(), "\n")
	return hits, err
}

func pathAndContentFromResult(result string) (path, content string) {
	x := strings.SplitAfterN(result, "md:", 2)
	if len(x) == 1 {
		x = strings.SplitAfterN(result, "md-", 2)
	}
	return x[0], x[1]
}

func sDateFromPath(path string) string {
	xPath := strings.Split(path, "/")
	sDate := xPath[len(xPath)-1]
	return sDate
}

func dateFromDateString(sDate string) (d time.Time, err error) {
	iYear, err := strconv.Atoi(sDate[:4])
	if err != nil {
		return d, err
	}
	iMonth, err := strconv.Atoi(sDate[4:6])
	if err != nil {
		return d, err
	}
	// time.Date constructor takes ints for every parameter except month
	// That's gotta be a time.Month, so we convert here
	month := time.Month(iMonth)
	iDay, err := strconv.Atoi(sDate[6:8])
	if err != nil {
		return d, err
	}
	d = time.Date(iYear, month, iDay, 0, 0, 0, 0, time.UTC)
	return d, err
}

func formatSearchResults(results []string) (err error) {
	// Start with a fresh empty newline
	fmt.Println()
	// Build array of timestamps (aka map keys) to sort later
	var dates timeSlice
	//c := color.New(color.FgCyan, color.Bold)
	// Problem: grep output is unsorted, we want results sorted by note timestamp
	// Solution: Build map of timestamps to search results
	entries := make(map[time.Time]string)
	for _, result := range results {
		if result != "" && result != "--" {
			// Split each grep result into filename (date) and output (note content)
			path, content := pathAndContentFromResult(result)
			sDate := sDateFromPath(path)
			d, err := dateFromDateString(sDate)
			if err != nil {
				return err
			}
			dates = updateTimeSlice(dates, d)
			entries = updateEntries(entries, d, content)
		}
	}
	// Sort array of timestamps
	sort.Sort(dates)
	// Print each search result in order of date logged
	for _, t := range dates {
		printHitsForDate(t, entries)
	}
	return err
}

func updateEntries(entries map[time.Time]string, d time.Time, content string) map[time.Time]string {
	if _, ok := entries[d]; ok {
		entries[d] += content + "\n"
	} else {
		entries[d] = content + "\n"
	}
	return entries
}

// Maintain timeSlice as a set
// TODO: Embed this logic in the type definition
// Or use a package implementation for sets
func updateTimeSlice(dates timeSlice, d time.Time) timeSlice {
	if !containsTime(dates, d) {
		dates = append(dates, d)
	}
	return dates
}

func printHitsForDate(t time.Time, hits map[time.Time]string) {
	y, m, d := t.Date()
	fmt.Println(y, m, d)
	fmt.Print(hits[t], "\n\n")
}

func containsTime(s []time.Time, t time.Time) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}
