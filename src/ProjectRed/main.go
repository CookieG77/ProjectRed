package main

import (
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	clearCmd()
	print(MenuRender())
	time.Sleep(1 / 24)
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
