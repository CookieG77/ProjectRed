package main

import (
	"PPR/InventoryTool"
	"os"
)

func main() {
	var crafts map[string]map[string]int
	InventoryTool.GetCraftList(&crafts, "data/crafts.json")

	var inv InventoryTool.Inventory = InventoryTool.InitInventory()
	var player map[string]interface{} = InventoryTool.InitPlayer()
	InventoryTool.AddGoldToPlayer(&player, 5)

	InventoryTool.AddItemToInventory(&inv, "O_WolfPelt", 2)
	InventoryTool.AddItemToInventory(&inv, "O_TrollSkin", 2)
	InventoryTool.AddItemToInventory(&inv, "O_Plume", 2)
	InventoryTool.AddItemToInventory(&inv, "O_BoarLeather", 2)
	InventoryTool.AddItemToInventory(&inv, "EC_Mage", 1)

	InventoryTool.Craft("EC_Adventurer", crafts, &inv, &player)

	InventoryTool.HurtPlayer(&player, 99)
	InventoryTool.EquipPlayerWith(&player, "EC_Adventurer", &inv)
	InventoryTool.PrintInventory(inv)
	InventoryTool.PrintPlayer(player)
	print("\n\n\n\n")
	InventoryTool.EquipPlayerWith(&player, "EC_Mage", &inv)
	InventoryTool.PrintInventory(inv)
	InventoryTool.PrintPlayer(player)

}

func getArgs() []string {
	return os.Args[1:]
}
