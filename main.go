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
		currentTime := time.Now()

		currentHour := fmt.Sprintf("%02d", currentTime.Hour())
		currentMin := fmt.Sprintf("%02d", currentTime.Minute())
		currentSec := fmt.Sprintf("%02d", currentTime.Second())

		// Format time and convert to Ascii art
		finalTime := fmt.Sprintf("%v:%v:%v", currentHour, currentMin, currentSec)
		asciiStr, _ := ascii.Render(finalTime)

		// Overwrite previous time
		lines := strings.Split(asciiStr, "\n")
		clearLines(len(lines))

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
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
