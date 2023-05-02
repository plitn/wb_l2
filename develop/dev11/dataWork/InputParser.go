package dataWork

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	// Define the date format you expect
	dateFormat := "2006-01-02"

	// Parse the date string using the expected format
	date, err := time.Parse(dateFormat, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	// Check that the parsed date is a valid date
	if date.IsZero() {
		return time.Time{}, fmt.Errorf("invalid date: %s", dateStr)
	}

	return date, nil
}
