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

	//Créations des données du joueur
	player := InventoryTool.InitPlayer()
	inv := InventoryTool.InitInventory()
	Scene.CreatePlayerWindow(classList, icons, &player, inv)

	//Execution du programme
	InventoryTool.AddItemToInventory(&inv, "EC_Mage", 1)
	InventoryTool.AddItemToInventory(&inv, "EC_Adventurer", 1)
	InventoryTool.EquipPlayerWith(&player, "EC_Mage", &inv, itemlist)
	player["mana"] = player["max_mana"].(int)
	InventoryTool.EquipPlayerWith(&player, "EC_Adventurer", &inv, itemlist)

	print(Smenu.SmenuRender(icons, bg, &player, itemlist, inv))

}
