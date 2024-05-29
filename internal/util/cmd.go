package util

import (
	"flag"
)

type Config struct {
	TimeFormat  string
	IsCentered  bool
	ShowVersion bool
	ShowDate    bool
}

func ParseFlags() *Config {
	var timeFormat string
	var isCentered bool
	var showVersion bool
	var showDate bool

	flag.StringVar(&timeFormat, "f", "24", "Time format (12 or 24)")
	flag.BoolVar(&isCentered, "C", false, "Center text")
	flag.BoolVar(&showVersion, "v", false, "Version")
	flag.BoolVar(&showDate, "d", false, "Show date")

	flag.Parse()

	return &Config{
		TimeFormat:  timeFormat,
		IsCentered:  isCentered,
		ShowVersion: showVersion,
		ShowDate:    showDate,
	}
}
