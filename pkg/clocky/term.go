package clocky

import (
	"os"
	"os/exec"
)

func ClearTerm() {
	cmdName := "clear"

	// Detect if on windows and run "cls" instead
	if os.Getenv("OS") == "Windows_NT" {
		cmdName = "cls"
	}

	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
