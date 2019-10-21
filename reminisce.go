package main

import "time"

func reminisce(configPath string) (err error) {
	now := time.Now()
	dates := make([]time.Time, 0)
	for i := -100; i < 0; i++ {
		dates = append(dates, now.AddDate(i, 0, 0))
	}
	dates = append(dates, now.AddDate(0, -6, 0))
	for _, date := range dates {
		f, err := initFile(date)
		if err != nil {
			return err
		}
		if Exists(f) {
			load(f, configPath)
		}
	}
	return err
}
