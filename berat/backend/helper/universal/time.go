package time

import (
	"errors"
	"time"
)

var errorIndoDateFormat = errors.New("date format must be dd-mm-yyyy")
var errorGlobalDateFormat = errors.New("date format must be yyyy-mm-dd")

func GetIndoFormattedDate(date string) (string, error) {
	layout := "2006-01-02T15:04:05Z"
	timeFormat, err := time.Parse(layout, date)
	if err != nil {
		return "", errorGlobalDateFormat
	}

	newLayout := timeFormat.Format("02-01-2006")
	return newLayout, nil
}

func GetGlobalFormattedDate(date string) (string, error) {
	layout := "02-01-2006"
	timeFormat, err := time.Parse(layout, date)
	if err != nil {
		return "", errorIndoDateFormat
	}

	newLayout := timeFormat.Format("2006-01-02")
	return newLayout, nil
}
