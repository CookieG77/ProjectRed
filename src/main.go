package main

import (
	"PPR/InventoryTool"
	"PPR/Scene"
	"PPR/Smenu"
	"image"
)

func main() {
	icons := make(map[string]image.Image)
	classList := make(map[string]map[string]interface{})
	monsterList := make(map[string]map[string]interface{})
	InventoryTool.LoadClassIcons(&icons, "ressource")
	InventoryTool.GetClassList(&classList, "data/classes.json")
	InventoryTool.GetClassList(&monsterList, "data/monsters.json")
	Scene.CreatePlayerWindow(classList, icons)
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
