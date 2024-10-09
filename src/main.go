package main

import (
	"PPR/InventoryTool"
	"PPR/Scene"
	"image"
	"os"
)

func main() {
	icons := make(map[string]image.Image)
	classList := make(map[string]map[string]interface{})
	monsterList := make(map[string]map[string]interface{})
	InventoryTool.LoadClassIcons(&icons, "ressource")
	InventoryTool.GetClassList(&classList, "data/classes.json")
	InventoryTool.GetClassList(&monsterList, "data/monsters.json")
	Scene.CreatePlayerWindow(classList, icons)
}

func getArgs() []string {
	return os.Args[1:]
}
