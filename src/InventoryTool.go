package PPR

import (
	"encoding/json"
	"fmt"
	"os"
)

type inventory = map[string]int

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

// |========================================================|
// |					Gestion Inventaire                  |
// |========================================================|

func InitInventory() inventory {
	return make(inventory)
}

func AddItemToInventory(inv *inventory, itemID string, quantity int) {
	//quantity doit être positif
	_, ok := (*inv)[itemID]
	if ok {
		(*inv)[itemID] += quantity
	} else {
		(*inv)[itemID] = quantity
	}
}

func RemoveItemFromInventory(inv *inventory, itemID string, quantity int) bool {
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

func PrintInventory(inv inventory) {
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
	res["maxhp"] = 0
	res["hp"] = 0
	res["gold"] = 0
	res["class"] = ""
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

// |========================================================|
// |				   	  Gestion Crafts                    |
// |========================================================|

func CanCraft(
	craftID string,
	craftlist map[string]map[string]int,
	inv inventory,
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
	inv inventory,
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
	inv *inventory,
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
