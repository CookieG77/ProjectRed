package main

import (
	"PPR"
	"os"
)

func main() {
	PPR.SmenuRender()
	// run := true
	// for !run {
	// 	clearCmd()

	// 	time.Sleep(time.Second / 24)
	// }

}

func getArgs() []string {
	return os.Args[1:]
}

// func clearCmd() {
// 	if runtime.GOOS == "windows" {
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	} else {
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	}
// }
