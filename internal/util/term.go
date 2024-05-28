package util

import (
	"fmt"
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
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func DisplayVersion() {
	fmt.Println("v1.2.4")
	os.Exit(0)
}
