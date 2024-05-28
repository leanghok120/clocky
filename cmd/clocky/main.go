package main

import (
	"fmt"
	"time"

	"github.com/leanghok120/clocky/internal/util"
	"github.com/mbndr/figlet4go"
)

func main() {
	cfg := util.ParseFlags()

	if cfg.ShowVersion {
		util.DisplayVersion()
	}

	ascii := figlet4go.NewAsciiRender()
	for {
		// Clear stdout
		util.ClearTerm()

		currentTime := util.GetTime(cfg.TimeFormat)
		asciiStr, _ := ascii.Render(currentTime)

		fmt.Print(asciiStr)
		time.Sleep(1 * time.Second)
	}
}
