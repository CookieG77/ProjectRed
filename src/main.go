package main

import (
	"PPR/InventoryTool"
	"PPR/Scene"
	"PPR/Smenu"
	"image"
)

func main() {
	//Importation et chargement des données et ressources
	icons := make(map[string]image.Image)
	InventoryTool.LoadClassIcons(&icons, "ressource")

	bg := make(map[string]image.Image)
	InventoryTool.LoadBG(&bg, "ressource")

	classList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&classList, "data/classes.json")

	monsterList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&monsterList, "data/monsters.json")

	//Créations des données du joueur
	player := InventoryTool.InitPlayer()
	//inv := InventoryTool.InitInventory()
	Scene.CreatePlayerWindow(classList, icons, &player)

	//Execution du programme

	print(Smenu.SmenuRender(icons, bg, &player))

}
