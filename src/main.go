package main

import (
	"PPR/Scene"
	"PPR/Smenu"
	"os"
)

func main() {
	// run := true
	// for !run {
	// 	clearCmd()

	// 	time.Sleep(time.Second / 24)
	// }
	Scene.CreatePlayerWindow()
	Smenu.SmenuRender()

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
