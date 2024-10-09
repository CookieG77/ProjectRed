package main

import (
	"PPR/InventoryTool"
	"PPR/Scene"
	"image"
)

func main() {
	//Importation et chargement des données et ressources
	icons := make(map[string]image.Image)
	InventoryTool.LoadClassIcons(&icons, "ressource")

	classList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&classList, "data/classes.json")

	monsterList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&monsterList, "data/monsters.json")

	//Créations des données du joueur
	player := InventoryTool.InitPlayer()
	//inv := InventoryTool.InitInventory()

	Scene.CreatePlayerWindow(classList, icons, &player)

	//Execution du programme

	// running := true
	// var scene int
	//print(Smenu.SmenuRender(icons))
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
