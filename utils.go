package main

import "time"

const (
	datefmt   = "01-02-06"
	outputfmt = "2006-01-02T15:04:05Z07:00"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func dateConverter(datestring string) time.Time {
	t, err := time.Parse(datefmt, datestring)
	check(err)

	return t
}
