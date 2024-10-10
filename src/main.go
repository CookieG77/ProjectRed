package main

import (
	"PPR/InventoryTool"
	"fmt"
	"image"

	"github.com/rivo/tview"
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

	craftList := make(map[string]map[string]int)
	InventoryTool.GetCraftList(&craftList, "data/crafts.json")

	//Créations des données du joueur
	player := InventoryTool.InitPlayer()
	inv := InventoryTool.InitInventory()

	//Execution du programme
	// Scene.CreatePlayerWindow(classList, class_icons, &player, inv)
	InventoryTool.AddItemToInventory(&inv, "O_WolfPelt", 2)
	InventoryTool.AddItemToInventory(&inv, "O_TrollSkin", 2)
	InventoryTool.AddItemToInventory(&inv, "O_BoarLeather", 2)

	InventoryTool.AddGoldToPlayer(&player, 100)
	// print(Smenu.SmenuRender(class_icons, bg, monster_icons, &player, itemlist, &inv, classList, skillList, monsterList))
	ShowForge(class_icons, bg, monster_icons, &player, itemlist, &inv, classList, skillList, monsterList, craftList)
}

func ShowForge(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	craftList map[string]map[string]int,
) {
	app := tview.NewApplication()
	list := tview.NewList()
	Craftable := InventoryTool.GetCraftableList(craftList, *inv, *player)
	a := 'a'
	for _, i := range Craftable {
		name := itemlist[string(i)]["name"].(string)
		list.AddItem(name, itemlist[string(i)]["description"].(string), a, func() {
			InventoryTool.Craft(i, craftList, inv, player)
			app.Stop()
			fmt.Println(inv)
		})
		a++
	}
	if err := app.SetRoot(list, true).SetFocus(list).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
