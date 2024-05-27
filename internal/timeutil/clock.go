package timeutil

import (
	"time"
)

func GetTime(format string) string {
	currentTime := time.Now()

	if format == "12" {
		return currentTime.Format("03:04:05 PM") // Returns the current time in HH:MM:SS AM/PM
	} else if format == "" || format == "24" {
		return currentTime.Format("15:04:05") // Returns the current time in HH:MM:SS
	}

	return "Invalid time format"
}
