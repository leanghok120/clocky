package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mbndr/figlet4go"
)

func main() {
	ascii := figlet4go.NewAsciiRender()
	var timeFormat string
	flag.StringVar(&timeFormat, "f", "24", "Time format")
	flag.Parse()

	// Clear stdout
	clearTerm()

	for {
		currentTime := getTime(timeFormat)
		asciiStr, _ := ascii.Render(currentTime)

		// Overwrite previous time
		lines := strings.Split(asciiStr, "\n")
		clearLines(len(lines))

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
}

func getTime(format string) string {
	currentTime := time.Now()

	if format == "12" {
		return currentTime.Format("03:04:05 PM")
	} else if format == "" || format == "24" {
		return currentTime.Format("15:04:05")
	}

	return "Invalid time format"
}

func clearLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[K")
	}
}

func clearTerm() {
	cmdName := "clear"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
