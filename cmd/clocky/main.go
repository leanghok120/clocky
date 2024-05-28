package main

import (
	"time"

	"github.com/leanghok120/clocky/internal/util"
	"github.com/mbndr/figlet4go"
	"github.com/rthornton128/goncurses"
)

func main() {
	cfg := util.ParseFlags()

	// Initialize ncurses
	stdscr, err := goncurses.Init()
	if err != nil {
		panic(err)
	}
	defer goncurses.End()

	if cfg.ShowVersion {
		util.DisplayVersion()
	}

	ascii := figlet4go.NewAsciiRender()
	for {
		// Clear stdscr
		stdscr.Clear()

		currentTime := util.GetTime(cfg.TimeFormat)
		asciiStr, _ := ascii.Render(currentTime)

		if cfg.IsCentered {
			util.CenterText(*stdscr, asciiStr)
		} else {
			stdscr.Print(asciiStr)
		}

		stdscr.Refresh()

		time.Sleep(1 * time.Second)
	}
}
