package game

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var Screen = bufio.NewWriter(os.Stdout)

func HideCursor() {
	fmt.Fprint(Screen, "\033[?25l")
}

func ShowCursor() {
	fmt.Fprint(Screen, "\033[?25h")
}

func MoveCursor(pos [2]int) {
	fmt.Fprintf(Screen, "\033[%d;%dH", pos[1], pos[0])
}

func Clear() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Craw(str string) {
	fmt.Fprint(Screen, str)
}

func Render() {
	Screen.Flush()
}
