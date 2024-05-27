package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/leanghok120/clocky/pkg/timeutil"
	"github.com/mbndr/figlet4go"
)

func main() {
	var timeFormat string
	flag.StringVar(&timeFormat, "f", "24", "Time format (12 or 24)")
	flag.Parse()

	ascii := figlet4go.NewAsciiRender()

	for {
		// Clear stdout
		timeutil.ClearTerm()

		currentTime := timeutil.GetTime(timeFormat)
		asciiStr, _ := ascii.Render(currentTime)

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
}
