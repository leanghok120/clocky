package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mbndr/figlet4go"
)

func main() {
	ascii := figlet4go.NewAsciiRender()

	// Clear stdout
	clearTerm()

	for {
		currentTime := getTime()
		asciiStr, _ := ascii.Render(currentTime)

		// Overwrite previous time
		lines := strings.Split(asciiStr, "\n")
		clearLines(len(lines))

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
}

func getTime() string {
	currentTime := time.Now()
	timeFormat := "24"

	if len(os.Args) >= 2 {
		timeFormat = os.Args[1]
	}

	if timeFormat == "12" {
		return currentTime.Format("03:04:05 PM")
	}

	return currentTime.Format("15:04:05")
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
