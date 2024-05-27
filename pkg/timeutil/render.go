package timeutil

import (
	"fmt"
	"strings"
)

func SplitLines(asciiStr string) []string {
	return strings.Split(asciiStr, "\n")
}

func ClearLines(n int) {
	for i := 0; i < n; i++ {
		fmt.Print("\033[1A")
		fmt.Print("\033[K")
	}
}
