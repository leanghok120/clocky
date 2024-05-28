package util

import (
	"flag"
)

type Config struct {
	TimeFormat string
	IsCentered bool
}

func ParseFlags() *Config {
	var timeFormat string
	var isCentered bool
	flag.StringVar(&timeFormat, "f", "24", "Time format (12 or 24)")
	flag.BoolVar(&isCentered, "C", false, "Center text")
	flag.Parse()

	return &Config{
		TimeFormat: timeFormat,
		IsCentered: isCentered,
	}
}
