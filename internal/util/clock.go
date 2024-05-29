package util

import (
	"time"

	"github.com/rthornton128/goncurses"
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

func GetDate() string {
	currentDate := time.Now()

	return currentDate.Format("02:01:2006") // Returns the date in DD:MM:YYYY
}

func PrintTime(cfg Config, stdscr goncurses.Window, str string) {
	if cfg.IsCentered {
		CenterText(stdscr, str)
	} else {
		stdscr.Print(str)
	}

	if cfg.ShowDate {
		currentDate := GetDate()

		stdscr.Print(currentDate)
	}
}
