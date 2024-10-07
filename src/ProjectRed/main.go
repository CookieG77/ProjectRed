package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	redColor := tcell.NewRGBColor(255, 0, 0)
	run := true
	box := tview.NewBox().SetBorder(true).SetTitle("[ RED_PROJECT_ULTIMATE ]")
	box.SetBorderColor(redColor)
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
	for !run {
		clearCmd()

		time.Sleep(time.Second / 24)
	}

}

func getArgs() []string {
	return os.Args[1:]
}

func clearCmd() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
