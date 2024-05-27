package clocky

import (
	"os"
	"os/exec"
)

func ClearTerm() {
	cmdName := "clear"
	cmd := exec.Command(cmdName)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
