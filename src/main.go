package main

import (
	"PPR/InventoryTool"
	"PPR/Scene"
	"PPR/Smenu"
	"image"
	"os"
	"slices"
)

func main() {
	//Importation et chargement des données et ressources
	class_icons := make(map[string]image.Image)
	InventoryTool.LoadClassIcons(&class_icons, "ressource")

	bg := make(map[string]image.Image)
	InventoryTool.LoadBG(&bg, "ressource")

	monster_icons := make(map[string]image.Image)
	InventoryTool.LoadMonsterIcons(&monster_icons, "ressource")

	itemlist := make(map[string]map[string]interface{})
	InventoryTool.GetItemList(&itemlist, "data/items.json")

	classList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&classList, "data/classes.json")

	monsterList := make(map[string]map[string]interface{})
	InventoryTool.GetClassList(&monsterList, "data/monsters.json")

	tradeList := make(map[string]map[string]map[string]int)
	InventoryTool.GetTradesList(&tradeList, "data/trades.json")

	skillList := make(map[string]map[string]interface{})
	InventoryTool.GetSkillList(&skillList, "data/skills.json")

	lootList := make(map[string]map[string]map[string][]int)
	InventoryTool.GetLootList(&lootList, "data/loots.json")

	craftList := make(map[string]map[string]int)
	InventoryTool.GetCraftList(&craftList, "data/crafts.json")

	donjonList := make(map[string]map[string]map[string]interface{})
	InventoryTool.GetDonjonList(&donjonList, "data/donjons.json")

	//Créations des données du joueur
	player := InventoryTool.InitPlayer()
	inv := InventoryTool.InitInventory()

	//Execution du programme
	go InventoryTool.PlayLoopMusic("ressource/music_background.mp3")
	Scene.CreatePlayerWindow(classList, class_icons, &player, inv)
	if slices.Contains(getArgs(), "-op") {
		InventoryTool.AddItemToInventory(&inv, "CP_Fireballbook", 1)
		InventoryTool.AddItemToInventory(&inv, "EC_Adventurer", 1)
		InventoryTool.AddItemToInventory(&inv, "EA_Adventurer", 1)
		InventoryTool.AddItemToInventory(&inv, "EL_Adventurer", 1)
		InventoryTool.AddItemToInventory(&inv, "EB_Adventurer", 1)
		InventoryTool.AddItemToInventory(&inv, "EC_Mage", 1)
		InventoryTool.AddItemToInventory(&inv, "EA_Mage", 1)
		InventoryTool.AddItemToInventory(&inv, "EL_Mage", 1)
		InventoryTool.AddItemToInventory(&inv, "EB_Mage", 1)
		InventoryTool.AddItemToInventory(&inv, "EA_Adventurer", 1)
		InventoryTool.AddItemToInventory(&inv, "EW_AdvancedMageStaff", 1)
		InventoryTool.AddItemToInventory(&inv, "EW_VampireDagger", 1)
		InventoryTool.AddItemToInventory(&inv, "CP_Heal", 2)
		InventoryTool.AddItemToInventory(&inv, "CP_Mana", 1)
		InventoryTool.AddItemToInventory(&inv, "CJ_Poison", 5)
		InventoryTool.PlayerLearnSkill(&player, "opskill")
		InventoryTool.PlayerLearnSkill(&player, "suicideskill")
		InventoryTool.AddItemToInventory(&inv, "O_WolfPelt", 2)
		InventoryTool.AddItemToInventory(&inv, "O_TrollSkin", 2)
		InventoryTool.AddItemToInventory(&inv, "O_BoarLeather", 2)
		InventoryTool.AddGoldToPlayer(&player, 1000)
		InventoryTool.EquipPlayerWith(&player, "EW_AdvancedMageStaff", &inv, itemlist)
		InventoryTool.EquipPlayerWith(&player, "EW_VampireDagger", &inv, itemlist)
	}
	if player["max_hp"].(int) > 0 {
		Smenu.SmenuRender(class_icons, bg, monster_icons, &player, itemlist, &inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonList)
	}

}

func getArgs() []string {
	return os.Args[1:]
}
