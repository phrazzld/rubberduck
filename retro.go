// rubberduck/retro.go

package main

import (
	"os"
	"time"
)

func retro(configPath string) (err error) {
	var dates timeSlice
	now := time.Now()
	for i := 7; i > 0; i-- {
		dates = append(dates, now.AddDate(0, 0, -i))
	}
	// Load last week's entries
	for _, date := range dates {
		f, err := initFile(date)
		if err != nil {
			return err
		}
		if Exists(f) {
			load(f, configPath)
		}
	}
	// Stamp today's entry with sprint retrospective prompts
	// And load it
	err = initRetroEntry(configPath)
	if err != nil {
		return err
	}
	f, err := initFile(now)
	if err != nil {
		return err
	}
	if Exists(f) {
		load(f, configPath)
	}
	return err
}

func initRetroEntry(configPath string) (err error) {
	// Build retro content
	retroContent := appendNewline("### Retrospective", 2)
	retroContent += appendNewline("#### What went well this week?", 2)
	retroContent += appendNewline("#### What did not go so well?", 2)
	retroContent += appendNewline("#### What did you learn?", 2)
	retroContent += appendNewline("#### What still puzzles you?", 2)
	retroContent += appendNewline("#### Analysis and Planning", 2)

	// Stamp day's rubberduck
	n := time.Now()
	d, t := initDatetime(n)
	f, err := initFile(n)
	if err != nil {
		return err
	}
	err = stamp(f, d, t, configPath)
	if err != nil {
		return err
	}

	// Append retro entry
	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(retroContent)
	if err != nil {
		return err
	}
	file.Sync()
	return err
}
