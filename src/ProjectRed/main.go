package main

import (
	"PPR"
	"os"
)

func main() {
	var crafts map[string]map[string]int
	PPR.GetCraftList(&crafts, "data/crafts.json")

	var inv map[string]int = PPR.InitInventory()
	var player map[string]interface{} = PPR.InitPlayer()
	PPR.AddGoldToPlayer(&player, 5)

	PPR.AddItemToInventory(&inv, "O_WolfPelt", 2)
	PPR.AddItemToInventory(&inv, "O_TrollSkin", 2)
	PPR.AddItemToInventory(&inv, "O_Plume", 2)
	PPR.AddItemToInventory(&inv, "O_BoarLeather", 2)

	PPR.Craft("EC_Adventurer", crafts, &inv, &player)
	PPR.PrintInventory(inv)
	PPR.PrintPlayer(player)

}

func getArgs() []string {
	return os.Args[1:]
}
