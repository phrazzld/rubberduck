// rubberduck/search.go

package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"
)

func search(duckPath string, searchTerm []string) []string {
	var b strings.Builder
	cmd := exec.Command("egrep", strings.Join(searchTerm, " "), "-R", duckPath, "--ignore-case", "-C", "1")
	cmd.Stdin = os.Stdin
	cmd.Stdout = &b
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	hits := strings.Split(b.String(), "\n")
	return hits
}

func formatSearchResults(results []string) {
	fmt.Println()
	c := color.New(color.FgCyan, color.Bold)
	var dateString string
	var year, day int
	var month time.Month
	var d time.Time
	var dates timeSlice
	entries := make(map[time.Time]string)
	for _, v := range results {
		if v != "" && v != "--" {
			var err error
			x := strings.SplitAfterN(v, "md:", 2)
			if len(x) == 1 {
				x = strings.SplitAfterN(v, "md-", 2)
			}
			path := strings.Split(x[0], "/")
			dateString = path[len(path)-1]
			year, err = strconv.Atoi(dateString[:4])
			monthAsInt, err := strconv.Atoi(dateString[4:6])
			if err != nil {
				fmt.Println(err)
			}
			month = time.Month(monthAsInt)
			day, err = strconv.Atoi(dateString[6:8])
			if err != nil {
				fmt.Println(err)
			}
			d = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
			if !containsTime(dates, d) {
				dates = append(dates, d)
			}
			year, month, day = d.Date()
			if _, ok := entries[d]; ok {
				entries[d] += x[1] + "\n"
			} else {
				entries[d] = x[1] + "\n"
			}
		}
	}
	sort.Sort(dates)
	for _, v := range dates {
		y, m, d := v.Date()
		c.Println(y, m, d)
		fmt.Print(entries[v], "\n\n")
	}
}

func containsTime(s []time.Time, t time.Time) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}
