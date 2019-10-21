// rubberduck/review.go

package main

import "time"

func review(configPath string) {
	var dates timeSlice
	now := time.Now()
	dates = append(dates, now.AddDate(0, -3, 0))
	dates = append(dates, now.AddDate(0, -1, 0))
	dates = append(dates, now.AddDate(0, 0, -7))
	dates = append(dates, now.AddDate(0, 0, -1))
	for _, date := range dates {
		f := initFile(date)
		if Exists(f) {
			load(f, configPath)
		}
	}
}
