package util

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rthornton128/goncurses"
)

func ClearTerm() {
	cmdName := "clear"

	// Detect if on windows and run "cls" instead
	if os.Getenv("OS") == "Windows_NT" {
		cmdName = "cls"
	}

	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func DisplayVersion() {
	fmt.Println("v1.3.3")
	os.Exit(0)
}

func CenterText(stdscr goncurses.Window, asciiStr string) {
	lines := strings.Split(asciiStr, "\n")
	maxY, maxX := stdscr.MaxYX()
	startY := (maxY - len(lines)) / 2
	for i, line := range lines {
		startX := (maxX - len(line)) / 2
		stdscr.MovePrint(startY+i, startX, line)
	}
}
