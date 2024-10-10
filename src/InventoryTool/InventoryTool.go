package InventoryTool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"
)

type Inventory = map[string]int

// |========================================================|
// |					Gestion data files                  |
// |========================================================|

// Ouvre le fichier json 'filepath' et le converti en variable lisible par golang dans 'res'
func GetCraftList(res *map[string]map[string]int, filepath string) bool {
	//Retourne 0 si aucune erreur n'a été rencontré lors du dépactage du json. Sinon retourne 1
	file, err1 := os.ReadFile(filepath)
	if err1 != nil || json.Unmarshal(file, res) != nil {
		return true
	}
	return false
}

// Ouvre le fichier json 'filepath' et le converti en variable lisible par golang dans 'res',
// et converti les nombres à virgules en nombres entier.
func GetItemList(res *map[string]map[string]interface{}, filepath string) bool {
	file, err1 := os.ReadFile(filepath)
	if err1 != nil || json.Unmarshal(file, res) != nil {
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

// Juste un alias de GetItemList pour simplifier la compréhension du code au main,
// car les fonctions marchent de la même manière
func GetClassList(res *map[string]map[string]interface{}, filepath string) bool {
	return GetItemList(res, filepath)
}

// Juste un alias de GetItemList pour simplifier la compréhension du code au main,
// car les fonctions marchent de la même manière
func GetMonsterList(res *map[string]map[string]interface{}, filepath string) bool {
	return GetItemList(res, filepath)
}

// Juste un alias de GetItemList pour simplifier la compréhension du code au main,
// car les fonctions marchent de la même manière
func GetSkillList(res *map[string]map[string]interface{}, filepath string) bool {
	return GetItemList(res, filepath)
}

// Ouvre le fichier json 'filepath' et le converti en variable lisible par golang dans 'res'
func GetTradesList(res *map[string]map[string]map[string]int, filepath string) bool {
	//Retourne 0 si aucune erreur n'a été rencontré lors du dépactage du json. Sinon retourne 1
	file, err1 := os.ReadFile(filepath)
	if err1 != nil || json.Unmarshal(file, res) != nil {
		return true
	}
	return false
}

// |========================================================|
// |					Gestion Inventaire                  |
// |========================================================|

// Initialise une variable de type inventaire.
func InitInventory() Inventory {
	return make(Inventory)
}

// Permet l'ajout de l'item portant l'id 'itemID'
// à l'inventaire 'inv' avec pour quantité 'quantity'.
func AddItemToInventory(inv *Inventory, itemID string, quantity int) {
	//quantity doit être positif
	_, ok := (*inv)[itemID]
	if ok {
		(*inv)[itemID] += quantity
	} else {
		(*inv)[itemID] = quantity
	}
}

// Permet la suppréssion de l'item portant l'id 'itemID'
// à l'inventaire 'inv' avec pour quantité 'quantity'.
// Renvoie 'true' si l'action est impossible (retirer plus d'item que le joueur en a)
// et sinon renvoie 'false'.
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

// Renvoie la liste des items commencant par les tags 'item_types'
// fournies dans l'inventaire 'inv'.
func getInventoryByItemType(inv Inventory, item_types []string) ([]string, []int) {
	res := []string{}
	res2 := []int{}
	for consumableID, quantity := range inv {
		for _, str := range item_types {
			if consumableID[:len(str)] == str && quantity > 0 {
				res = append(res, consumableID)
				res2 = append(res2, quantity)
			}
		}
	}
	return res, res2
}

// Renvoie la liste des items ne commencant pas par les tags 'item_types'
// fournies dans l'inventaire 'inv'.
func getInventoryNotItemType(inv Inventory, item_types []string) ([]string, []int) {
	res := []string{}
	res2 := []int{}
	for consumableID, quantity := range inv {
		ok := true
		for _, str := range item_types {
			if consumableID[:len(str)] == str {
				ok = false
				break
			}
		}
		if ok && quantity > 0 {
			res = append(res, consumableID)
			res2 = append(res2, quantity)
		}
	}
	return res, res2
}

// Renvoie la liste des items consomables
// ainsi qu'une autre liste contenant la quantité de chaques chacun.
func GetInventoryConsumables(inv Inventory) ([]string, []int) {
	return getInventoryByItemType(inv, []string{"CP"})
}

// Renvoie la liste des items équipables
// ainsi qu'une autre liste contenant la quantité de chacun.
func GetInventoryEquipable(inv Inventory) ([]string, []int) {
	return getInventoryByItemType(inv, []string{"EC", "EA", "EL", "EB", "EW"})
}

// Renvoie la liste des autres items
func GetInventoryOther(inv Inventory) ([]string, []int) {
	return getInventoryNotItemType(inv, []string{"EC", "EA", "EL", "EB", "EW", "CP"})
}

// Fonction à but de test uniquement, permet l'affichage de 'inv' sous la forme d'un dictionnaire.
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

// Fonction permettant l'initialisation des variables d'un joueur.
func InitPlayer() map[string]interface{} {
	res := make(map[string]interface{})
	res["name"] = ""
	res["class"] = ""
	res["max_hp"] = 0
	res["hp"] = 0
	res["max_mana"] = 0
	res["mana"] = 0
	res["gold"] = 0
	res["EquipmentHead"] = ""
	res["EquipmentTorso"] = ""
	res["EquipmentLegs"] = ""
	res["EquipmentBoots"] = ""
	res["EquipmentWeapon"] = ""
	res["skills"] = []string{}
	res["lvl"] = 0
	return res
}

// Fonction à but de test uniquement.
// Permet l'affichage de 'player' sous la forme d'un dictionnaire.
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

// Fonction permettant l'ajout de pièce d'or à un joueur.
func AddGoldToPlayer(player *map[string]interface{}, quantity int) {
	(*player)["gold"] = (*player)["gold"].(int) + quantity
}

// Permet la supréssion d'or au joueur 'player' avec pour quantité 'quantity'.
// Renvoie 'true' si l'action est impossible (retirer plus d'or que le joueur en a)
// et sinon renvoie 'false'.
func RemoveGoldFromPlayer(player *map[string]interface{}, quantity int) bool {
	val := (*player)["gold"].(int)
	if val < quantity {
		return true
	}
	(*player)["gold"] = (*player)["gold"].(int) - quantity
	return false
}

// Permet de savoir si oui ou non le joueur est mort (quand il n'a plus de PV).
func IsPlayerDead(player map[string]interface{}) bool {
	return player["hp"].(int) <= 0
}

func HealPlayer(player *map[string]interface{}, quantity int) {
	tmp := (*player)["hp"].(int) + quantity
	if tmp > (*player)["max_hp"].(int) {
		(*player)["hp"] = (*player)["max_hp"]
	} else {
		(*player)["hp"] = tmp
	}
}

// Permet d'infliger 'quantity' dégat(s) au joueur 'player'.
// Si l'on inflige plus de dégats que de PV restant au joueur, ses PV sont mits à 0.
// Renvoie true si le joueur meurt des dégats sinon renvoie false.
func HurtPlayer(player *map[string]interface{}, quantity int) bool {
	if (*player)["hp"].(int)-quantity <= 0 {
		(*player)["hp"] = 0
		return true
	}
	(*player)["hp"] = (*player)["hp"].(int) - quantity
	return false
}

// Permet de remplir le mana du joueur 'player' de 'quantity' mana.
func HealPlayerMana(player *map[string]interface{}, quantity int) {
	tmp := (*player)["mana"].(int) + quantity
	if tmp > (*player)["max_mana"].(int) {
		(*player)["mana"] = (*player)["max_mana"]
	} else {
		(*player)["mana"] = tmp
	}
}

// Permet de consomé 'quantity' mana au joueur 'player'.
// Si l'on consomme plus de mana que de mana restant au joueur, son mana est mit à 0.
// Renvoie true si le joueur n'a plus des dégats sinon renvoie false.
func UsePlayerMana(player *map[string]interface{}, quantity int) bool {
	if (*player)["mana"].(int)-quantity <= 0 {
		(*player)["mana"] = 0
		return true
	}
	(*player)["mana"] = (*player)["mana"].(int) - quantity
	return false
}

// Permet d'infliger 'quantity' dégats au joueur 'player' chaques secondes,
// pendant 'duration' secondes.
func HurtPlayerDPS(player *map[string]interface{}, quantity int, duration int) {
	if duration > 0 || quantity > 0 {
		time.Sleep(1 * time.Second)
		HurtPlayer(player, quantity)
		go HurtPlayerDPS(player, quantity, duration-1)
	}
}

// Renvoie 'true' si le joueur 'player' peut équipé un item 'equipment',
// en vérifiant si il en possède au moins un dans l'inventaire 'inv'.
// Sinon renvoie 'false'.
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

// Permet d'jouter le skill 'skill' à la liste des skills connu du joueur 'player'
func PlayerLearnSkill(player *map[string]interface{}, skill string) {
	(*player)["skills"] = append((*player)["skills"].([]string), skill)
}

// Equipe le joueur 'player' avec l'équipement 'itemID'.
// Déséquipe l'équipement actuelle dans le même emplacement si il en a déjà un.
// Les items équipé et déséquipé sont ajouté et retirer de l'inventaire 'inv'.
// Cette fonction à pour objectif d'être utiliser après
// avoir vérifier si l'item 'itemID' peut être équipé.
// Applique aussi les effets des armures au joueur comme "maxhealth" et "maxmana".
func EquipPlayerWith(
	player *map[string]interface{},
	itemID string, inv *Inventory,
	itemList map[string]map[string]interface{},
) {
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
	// Si le slot n'est pas vide,
	// on rend l'équipement au joueur dans son inventaire puis on recalcule les stats.
	if (*player)[slot].(string) != "" {
		AddItemToInventory(inv, (*player)[slot].(string), 1)
		switch itemList[(*player)[slot].(string)]["type"].(string) {
		case "maxhealth":
			{
				(*player)["max_hp"] = (*player)["max_hp"].(int) - itemList[(*player)[slot].(string)]["value"].(int)
				if (*player)["hp"].(int) > (*player)["max_hp"].(int) {
					(*player)["hp"] = (*player)["max_hp"].(int)
				}
			}
		case "maxmana":
			{
				(*player)["max_mana"] = (*player)["max_mana"].(int) - itemList[(*player)[slot].(string)]["value"].(int)
				if (*player)["mana"].(int) > (*player)["max_mana"].(int) {
					(*player)["mana"] = (*player)["max_mana"].(int)
				}
			}
		}
	}
	//On équipe le joueur avec 'itemID' dans le bon slot.
	(*player)[slot] = itemID
	RemoveItemFromInventory(inv, itemID, 1)
	switch itemList[itemID]["type"].(string) {
	case "maxhealth":
		(*player)["max_hp"] = (*player)["max_hp"].(int) + itemList[itemID]["value"].(int)
	case "maxmana":
		(*player)["max_mana"] = (*player)["max_mana"].(int) + itemList[itemID]["value"].(int)
	}

}

// |========================================================|
// |				   	  Gestion Crafts                    |
// |========================================================|

// Fonction pour vérifier si le joueur 'player' peut crafter l'objet 'craftID'
// (vérifie les matériaux et l'or nécessaire au craft).
// Renvoie 'true' si le craft est possible sinon renvoie 'false'.
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

// Fonction renvoyant un slice (une liste),
// contenant l'id des objets que le joueur 'player' peut craft
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

// Craft l'item 'itemID' et retire les matériaux nécessaire de l'inventaire 'inv'
// du joueur 'player' ainsi que la quantité d'or nécessaire
// Cette fonction à pour objectif d'être utiliser après
// avoir vérifier si l'item 'itemID' peut être craft.
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

// |========================================================|
// |	     Outils pour la création du personnage          |
// |========================================================|

// Renvoie un text comprennant la déscription de toutes les classes dans 'listclass'
func MakeClassDescs(listclass map[string]map[string]interface{}) string {
	res := ""
	for _, v := range listclass {
		res += v["stats_desc"].(string) + "\n"
	}
	return res
}

// Renvoie la liste du nom de toutes les classes dans 'listclass'
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

// Charge les images commencant par 'icon_' placé dans le dossier 'filepath',
// puis les range dans 'imglst'.Si une erreur apparai durant le chargement des images,
// la fonction s'arrète et renvoie 'true', sinon renvoie 'false'.
func LoadClassIcons(imglst *map[string]image.Image, filepath string) bool {
	files, err := os.ReadDir(filepath)
	if err != nil {
		return true
	}
	for _, fname := range files {
		if (fname.Name())[:4] == "icon" {
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

func LoadBG(imglst *map[string]image.Image, filepath string) bool {
	files, err := os.ReadDir(filepath)
	if err != nil {
		return true
	}
	for _, fname := range files {
		if (fname.Name())[:2] == "bg" {
			tmp, err2 := TViewMakeImg(filepath + "/" + fname.Name())
			if err2 {
				return true
			}
			name := fname.Name()[3 : len(fname.Name())-4]
			(*imglst)[name] = tmp
		}
	}
	return false
}

func LoadMonsterIcons(imglst *map[string]image.Image, filepath string) bool {
	files, err := os.ReadDir(filepath)
	if err != nil {
		return true
	}
	for _, fname := range files {
		if (fname.Name())[:5] == "micon" {
			tmp, err2 := TViewMakeImg(filepath + "/" + fname.Name())
			if err2 {
				return true
			}
			name := fname.Name()[6 : len(fname.Name())-4]
			(*imglst)[name] = tmp
		}
	}
	return false
}

// Fonctions permettant de chargé et décodé une image en 'png',
// pour utilisation avec la librairie Tview.
func TViewMakeImg(addresse string) (image.Image, bool) {
	IMGbyte, err := os.ReadFile(addresse)
	graphics, err2 := png.Decode(bytes.NewReader(IMGbyte))
	if err2 != nil || err != nil {
		return nil, true
	}
	return graphics, false
}

// |========================================================|
// |	              Gestion pnj marchant                  |
// |========================================================|

// Renvoie 'true' si le joueur 'player' peut acheter l'item 'itemID' auprès du vendeur 'trader'.
// Sinon renvoie 'false'. (Qu'il n'ai pas assez d'or où que l'item n'y soit pas vendu)
func CanBuyItemFrom(
	player map[string]interface{},
	tradeList map[string]map[string]map[string]int,
	trader string,
	itemID string,
) bool {
	val, ok := tradeList[trader]["sells"][itemID]
	return ok && player["gold"].(int) >= val
}

// Renvoie 'true' si le vendeur 'trader' peut acheter l'item 'itemID' auprès du joueur 'player'.
// Sinon renvoie 'false'.(Qu'il n'ai pas l'item sur lui ou que le vendeur n'achaite pas cette item)
func CanSellItemTo(
	inv map[string]int,
	tradeList map[string]map[string]map[string]int,
	trader string,
	itemID string,
) bool {
	_, ok := tradeList[trader]["buys"][itemID]
	return ok && inv[itemID] > 0
}

// Permet l'achat de l'item 'itemID' auprès du vendeur 'trader' par le joueur 'player'.
func BuyItemFrom(
	player *map[string]interface{},
	inv *map[string]int,
	tradeList map[string]map[string]map[string]int,
	trader string,
	itemID string,
) {
	RemoveGoldFromPlayer(player, tradeList[trader]["sells"][itemID])
	AddItemToInventory(inv, itemID, 1)
}

// Permet la vente de l'item 'itemID' auprès du vendeur 'trader' par le joueur 'player'
func SellItemTo(
	player *map[string]interface{},
	inv *map[string]int,
	tradeList map[string]map[string]map[string]int,
	trader string,
	itemID string,
) {
	AddGoldToPlayer(player, tradeList[trader]["buys"][itemID])
	RemoveItemFromInventory(inv, itemID, 1)
}
