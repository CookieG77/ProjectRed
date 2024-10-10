package combattool

import (
	"PPR/InventoryTool"
	"math/rand/v2"
	"strconv"
	"time"
)

func CanPlayerUseSkill(
	player map[string]interface{},
	skill string,
	skillList map[string]map[string]interface{},
) bool {
	return player["mana"].(int) >= skillList[skill]["mana_cost"].(int)
}

func GetUseableSkills(
	player map[string]interface{},
	skillList map[string]map[string]interface{},
) []string {
	res := []string{}
	for k := range skillList {
		if CanPlayerUseSkill(player, k, skillList) {
			res = append(res, k)
		}
	}
	return res
}

func UseSkill(
	player *map[string]interface{},
	monster *map[string]interface{},
	skill string,
	skillList map[string]map[string]interface{},
) {
	switch skillList[skill]["type"].(string) {
	case "dmg":
		{
			if skillList[skill]["target_player"].(bool) {
				InventoryTool.HurtPlayer(player, skillList[skill]["atk_points"].(int))
				InventoryTool.UsePlayerMana(player, skillList[skill]["mana_cost"].(int))
			} else {
				HurtMonster(monster, skillList[skill]["atk_points"].(int))
				InventoryTool.UsePlayerMana(player, skillList[skill]["mana_cost"].(int))
			}
		}
	case "heal":
		{
			if skillList[skill]["target_player"].(bool) {
				InventoryTool.HealPlayer(player, skillList[skill]["atk_points"].(int))
				InventoryTool.UsePlayerMana(player, skillList[skill]["mana_cost"].(int))
			} else {
				HealMonster(monster, skillList[skill]["atk_points"].(int))
				InventoryTool.UsePlayerMana(player, skillList[skill]["mana_cost"].(int))
			}
		}

	}
}

func UseConsumable(
	player *map[string]interface{},
	monster *map[string]interface{},
	consumableID string,
	itemList map[string]map[string]interface{},
	inv *map[string]int,
) {
	monsterPlayer := consumableID[:2] == "CP"
	itemData := itemList[consumableID]
	switch itemData["type"].(string) {
	case "heal":
		{
			if monsterPlayer {
				InventoryTool.HealPlayer(player, itemData["value"].(int))
			} else {
				HealMonster(monster, itemData["value"].(int))
			}
		}
	case "manaheal":
		{
			if monsterPlayer {
				InventoryTool.HealPlayerMana(player, itemData["value"].(int))
			} else {
				HealMonsterMana(monster, itemData["value"].(int))
			}
		}
	case "dps":
		{
			if monsterPlayer {
				InventoryTool.HurtPlayerDPS(player, itemData["value"].(int), itemData["duration"].(int))
			} else {
				HurtMonserDPS(monster, itemData["value"].(int), itemData["duration"].(int))
			}
		}
	case "spellbook":
		{
			if monsterPlayer {
				InventoryTool.PlayerLearnSkill(player, itemData["value"].(string))
			} else {
				monsterLearnSkill(monster, itemData["value"].(string))
			}
		}
	}
	InventoryTool.RemoveItemFromInventory(inv, consumableID, 1)
}

func MonsterAttack(
	player *map[string]interface{},
	monster *map[string]interface{},
	monsterList map[string]map[string]interface{},
	turn int,
) (string, bool) {
	res := ""
	switch (*monster)["spe"].(string) {
	case "crit":
		{
			if turn != 0 && turn%(*monster)["tour"].(int) == 0 {
				InventoryTool.HurtPlayer(player, (*monster)["atk_points"].(int)*(*monster)["special"].(int))
				res += (*monster)["atk_msg"].(string) + " Vous subbissez " + strconv.Itoa((*monster)["atk_points"].(int)) + " dégats." + (*monster)["spe_msg"].(string)
			} else {
				InventoryTool.HurtPlayer(player, (*monster)["atk_points"].(int))
				res += (*monster)["atk_msg"].(string)
			}
		}
	case "heal":
		{
			InventoryTool.HurtPlayer(player, (*monster)["atk_points"].(int))
			res += (*monster)["atk_msg"].(string) + " Vous subbissez " + strconv.Itoa((*monster)["atk_points"].(int)) + " dégats."
			if turn != 0 && turn%(*monster)["tour"].(int) == 0 {
				HealMonster(monster, (*monster)["special"].(int))
				res += (*monster)["spe_msg"].(string)
			}
		}
	case "steal_and_run":
		{
			if turn != 0 && turn%(*monster)["tour"].(int) == 0 {
				InventoryTool.RemoveGoldFromPlayer(player, (*monster)["special"].(int))
				res += (*monster)["spe_msg"].(string)
				return res, true
			}
			res += (*monster)["atk_msg"].(string) + " Vous subbissez " + strconv.Itoa((*monster)["atk_points"].(int)) + " dégats."
			InventoryTool.HurtPlayer(player, (*monster)["atk_points"].(int))
		}
	case "dpt":
		{
			res += (*monster)["atk_msg"].(string) + " Vous subbissez " + strconv.Itoa((*monster)["atk_points"].(int)) + " dégats."
			if (*monster)["tour"].(int) < rand.IntN(100) {
				InventoryTool.HurtPlayer(player, (*monster)["special"].(int))
				res += (*monster)["spe_msg"].(string)
			}
			InventoryTool.HurtPlayer(player, (*monster)["atk_points"].(int))
		}
	case "reduce_dmg":
		{
			res += (*monster)["atk_msg"].(string) + " Vous subbissez " + strconv.Itoa((*monster)["atk_points"].(int)) + " dégats."
			InventoryTool.HurtPlayer(player, (*monster)["atk_points"].(int))
		}
	}
	return res, false
}

func IsMonsterDead(monster map[string]interface{}) bool {
	return monster["hp"].(int) <= 0
}

func HealMonster(monster *map[string]interface{}, quantity int) {
	tmp := (*monster)["hp"].(int) + quantity
	if tmp > (*monster)["max_hp"].(int) {
		(*monster)["hp"] = (*monster)["max_hp"]
	} else {
		(*monster)["hp"] = tmp
	}
}

func HurtMonster(monster *map[string]interface{}, quantity int) bool {
	realquantity := quantity
	if (*monster)["spe"].(string) == "reduce_dmg" {
		realquantity -= (*monster)["special"].(int)
		if realquantity < 0 {
			realquantity = 0
		}
	}
	if (*monster)["hp"].(int)-realquantity <= 0 {
		(*monster)["hp"] = 0
		return true
	}
	(*monster)["hp"] = (*monster)["hp"].(int) - realquantity
	return false
}

func HealMonsterMana(monster *map[string]interface{}, quantity int) {
	tmp := (*monster)["mana"].(int) + quantity
	if tmp > (*monster)["max_mana"].(int) {
		(*monster)["mana"] = (*monster)["max_mana"]
	} else {
		(*monster)["mana"] = tmp
	}
}

func UseMonserMana(monster *map[string]interface{}, quantity int) bool {
	if (*monster)["mana"].(int)-quantity <= 0 {
		(*monster)["mana"] = 0
		return true
	}
	(*monster)["mana"] = (*monster)["mana"].(int) - quantity
	return false
}

func HurtMonserDPS(monster *map[string]interface{}, quantity int, duration int) {
	if duration > 0 {
		time.Sleep(1 * time.Second)
		HurtMonster(monster, quantity)
		go HurtMonserDPS(monster, quantity, duration-1)
	}
}

func monsterLearnSkill(monster *map[string]interface{}, skill string) {
	(*monster)["skills"] = append((*monster)["skills"].([]string), skill)
}

func GenRandMonster(monsterList map[string]map[string]interface{}) map[string]interface{} {
	monster := make(map[string]interface{})
	randomM := rand.IntN(len(monsterList))
	i := 0
	for k, v := range monsterList {
		if i == randomM {
			monster = v
			monster["id"] = k
			monster["hp"] = monster["max_hp"].(int)
			break
		}
		i++
	}
	return monster
}
