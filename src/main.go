package main

import (
	"PPR/Scene"
	"PPR/Smenu"
	"os"
)

func main() {
	// idk like If no player save somewhere skip
	Scene.CreatePlayerWindow()

	// running := true
	// var scene int
	print(Smenu.SmenuRender())
	// for running {
	// 	scene = Smenu.SmenuRender()
	// 	// switching beetween scenes
	// 	switch scene {
	// 	// Village
	// 	case 1:
	// 		fmt.Println("one")
	// 	// Foret
	// 	case 2:
	// 		fmt.Println("two")
	// 	// Inventaire
	// 	case 3:
	// 		fmt.Println("three")
	// 	// QUIT
	// 	case 4:
	// 		// idk os . cmd do Control+C for the user
	// 	// ERROR
	// 	case 5:
	// 		// idk something like, write in error.txt error has happened
	// 	}

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
