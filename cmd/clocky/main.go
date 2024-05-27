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

	// Clear stdout
	timeutil.ClearTerm()

	ascii := figlet4go.NewAsciiRender()
	for {
		currentTime := timeutil.GetTime(timeFormat)
		asciiStr, _ := ascii.Render(currentTime)

		// Overwrite previous time
		lines := timeutil.SplitLines(asciiStr)
		timeutil.ClearLines(len(lines))

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
}
