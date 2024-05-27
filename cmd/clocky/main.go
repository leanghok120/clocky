package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/leanghok120/clocky/pkg/clocky"
	"github.com/mbndr/figlet4go"
)

func main() {
	var timeFormat string
	flag.StringVar(&timeFormat, "f", "24", "Time format (12 or 24)")
	flag.Parse()

	// Clear stdout
	clocky.ClearTerm()

	ascii := figlet4go.NewAsciiRender()
	for {
		currentTime := clocky.GetTime(timeFormat)
		asciiStr, _ := ascii.Render(currentTime)

		// Overwrite previous time
		lines := clocky.SplitLines(asciiStr)
		clocky.ClearLines(len(lines))

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
}
