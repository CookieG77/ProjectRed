package InventoryTool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"os"
)

type Inventory = map[string]int

// |========================================================|
// |					Gestion data files                  |
// |========================================================|

func OpenJson(res *map[string]map[string]interface{}, filepath string) bool {
	//Retourne 0 si aucune erreur n'a été rencontré lors du dépactage du json. Sinon retourne 1
	file, err1 := os.ReadFile(filepath)
	if err1 != nil || json.Unmarshal(file, res) != nil {
		return true
	}
	return false
}

func GetCraftList(res *map[string]map[string]int, filepath string) bool {
	//Retourne 0 si aucune erreur n'a été rencontré lors du dépactage du json. Sinon retourne 1
	file, err1 := os.ReadFile(filepath)
	if err1 != nil || json.Unmarshal(file, res) != nil {
		return true
	}
	return false
}

func InitItemList(res *map[string]map[string]interface{}, filepath string) bool {
	if OpenJson(res, filepath) {
		return true
	}
	for k, v := range *res {
		ninter := make(map[string]interface{})
		for kk, vv := range v {
			tmp, ok := vv.(float64)
			if ok {
				ninter[kk] = int(tmp)
			} else {
				ninter[kk] = vv
			}
		}
		(*res)[k] = ninter
	}
	return false
}

func GetClassList(res *map[string]map[string]interface{}, filepath string) bool {
	return InitItemList(res, filepath)
}

func GetMonsterList(res *map[string]map[string]interface{}, filepath string) bool {
	return InitItemList(res, filepath)
}

// |========================================================|
// |					Gestion Inventaire                  |
// |========================================================|

func InitInventory() Inventory {
	return make(Inventory)
}

func AddItemToInventory(inv *Inventory, itemID string, quantity int) {
	//quantity doit être positif
	_, ok := (*inv)[itemID]
	if ok {
		(*inv)[itemID] += quantity
	} else {
		(*inv)[itemID] = quantity
	}
}

func RemoveItemFromInventory(inv *Inventory, itemID string, quantity int) bool {
	//quantity doit être positif
	//Retourne false si on a réussit à lui retirer sinon renvoie true.
	invquantity, ok := (*inv)[itemID]
	if ok {
		if invquantity < quantity {
			return true
		}
		(*inv)[itemID] -= quantity
	} else {
		return true
	}
	return false
}

func PrintInventory(inv Inventory) {
	print("{\n")
	for k, v := range inv {
		fmt.Printf(string('\u0009')+"%s : %d\n", k, v)
	}
	print("}\n")
}

// |========================================================|
// |					  Gestion Joueur                    |
// |========================================================|

func InitPlayer() map[string]interface{} {
	res := make(map[string]interface{})
	res["name"] = ""
	res["maxhp"] = 100
	res["hp"] = 100
	res["mana"] = 0
	res["maxmana"] = 0
	res["gold"] = 0
	res["class"] = ""
	res["EquipmentHead"] = ""
	res["EquipmentTorso"] = ""
	res["EquipmentLegs"] = ""
	res["EquipmentBoots"] = ""
	res["EquipmentWeapon"] = ""
	return res
}

func PrintPlayer(player map[string]interface{}) {
	print("{\n")
	for k, v := range player {
		_, ok := v.(int)
		if ok {
			fmt.Printf(string('\u0009')+"%s : %d\n", k, v)
		} else {
			fmt.Printf(string('\u0009')+"%s : %s\n", k, v)
		}
	}
	print("}\n")
}

func AddGoldToPlayer(player *map[string]interface{}, quantity int) {
	(*player)["gold"] = (*player)["gold"].(int) + quantity
}

func RemoveGoldFromPlayer(player *map[string]interface{}, quantity int) bool {
	val := (*player)["gold"].(int)
	if val < quantity {
		return true
	}
	(*player)["gold"] = (*player)["gold"].(int) - quantity
	return false
}

func IsPlayerDead(player map[string]interface{}) bool {
	return player["pv"].(int) <= 0
}

func HurtPlayer(player *map[string]interface{}, quantity int) bool {
	if (*player)["hp"].(int)-quantity <= 0 {
		(*player)["hp"] = 0
		return true
	}
	(*player)["hp"] = (*player)["hp"].(int) - quantity
	return false
}

func PlayerCanEquip(
	player map[string]interface{},
	equipment string,
	inv Inventory,
) bool {
	if val, ok := inv[equipment]; ok && val > 0 {
		return true
	}
	return false
}

func EquipPlayerWith(player *map[string]interface{}, itemID string, inv *Inventory) {
	slot := ""
	switch itemID[:2] {
	case "EC":
		slot = "EquipmentHead"

	case "EA":
		slot = "EquipmentTorso"

	case "EL":
		slot = "EquipmentLegs"

	case "EB":
		slot = "EquipmentBoots"

	case "W":
		slot = "EquipmentWeapon"
	}
	if (*player)[slot] != "" {
		AddItemToInventory(inv, (*player)[slot].(string), 1)
	}
	(*player)[slot] = itemID
	RemoveItemFromInventory(inv, itemID, 1)
}

// |========================================================|
// |				   	  Gestion Crafts                    |
// |========================================================|

func CanCraft(
	craftID string,
	craftlist map[string]map[string]int,
	inv Inventory,
	player map[string]interface{},
) bool {
	craft := craftlist[craftID]
	for item, quantity := range craft {
		if item == "cvalue" {
			d, ok := player["gold"].(int)
			if ok && int(d) < quantity {
				return false
			}
		} else {
			storedquantity, ok := inv[item]
			if !ok || storedquantity < quantity {
				return false
			}
		}
	}
	return true
}

func GetCraftableList(
	craftlist map[string]map[string]int,
	inv Inventory,
	player map[string]interface{},
) []string {
	var res []string
	for k := range craftlist {
		if CanCraft(k, craftlist, inv, player) {
			res = append(res, k)
		}
	}
	return res
}

func Craft(
	craftID string,
	craftlist map[string]map[string]int,
	inv *Inventory,
	player *map[string]interface{},
) {
	for k, v := range craftlist[craftID] {
		if k == "cvalue" {
			RemoveGoldFromPlayer(player, v)
		} else {
			RemoveItemFromInventory(inv, k, v)
		}
	}
	AddItemToInventory(inv, craftID, 1)
}

//
//
//

func MakeClassDescs(listclass map[string]map[string]interface{}) string {
	res := ""
	for _, v := range listclass {
		res += v["stats_desc"].(string) + "\n"
	}
	return res
}

func GetClassNames(listclass map[string]map[string]interface{}) []string {
	res := []string{}
	for _, v := range listclass {
		res = append(res, v["class"].(string))
	}
	return res
}

// |========================================================|
// |				 Gestion ressource files                |
// |========================================================|

func LoadClassIcons(imglst *map[string]image.Image, filepath string) bool {
	files, err := os.ReadDir(filepath)
	if err != nil {
		return true
	}
	for _, fname := range files {
		print(fname.Name() + "\n")
		if (fname.Name())[:4] == "icon" {
			print(filepath + "/" + fname.Name())
			tmp, err2 := TViewMakeImg(filepath + "/" + fname.Name())
			if err2 {
				return true
			}
			name := fname.Name()[5 : len(fname.Name())-4]
			(*imglst)[name] = tmp
		}
	}
	return false
}

func TViewMakeImg(addresse string) (image.Image, bool) {
	IMGbyte, err := os.ReadFile(addresse)
	graphics, err2 := png.Decode(bytes.NewReader(IMGbyte))
	if err2 != nil || err != nil {
		return nil, true
	}
	return graphics, false
}
