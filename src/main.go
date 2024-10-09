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

	itemlist := make(map[string]map[string]interface{})
	InventoryTool.GetItemList(&itemlist, "data/items.json")

	classList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&classList, "data/classes.json")

	monsterList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&monsterList, "data/monsters.json")

	tradeList := make(map[string]map[string]map[string]int)
	InventoryTool.GetTradesList(&tradeList, "data/trades.json")
	//Créations des données du joueur
	player := InventoryTool.InitPlayer()
	inv := InventoryTool.InitInventory()

	//Execution du programme
	Scene.CreatePlayerWindow(classList, icons, &player, inv)
	print(Smenu.SmenuRender(icons, bg, &player, itemlist, inv))

}
