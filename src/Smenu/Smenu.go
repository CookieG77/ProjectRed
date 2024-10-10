package Smenu

import (
	"image"
	"strconv"
	"time"

	combattool "PPR/CombatTool"
	"PPR/InventoryTool"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SmenuRender(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
) int {
	sceneValue := 0

	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	// =============================== Partie gauche ====================================

	// ++++++ header ++++++
	headergauche := tview.NewTextView().SetText("Intitié le voyage en direction de :")
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	// ++++++ images ++++++
	imageV := tview.NewImage()
	imageV.SetBorder(true)
	imageV.SetImage(bg_imgs["village_short"])

	imageF := tview.NewImage()
	imageF.SetBorder(true)
	imageF.SetImage(bg_imgs["forest_short"])

	// ++++++ Boutons ++++++
	buttonF := tview.NewButton("go to forest").SetSelectedFunc(func() {
		sceneValue = 2
		app.Stop()
		Monster := combattool.GenRandMonster(monsterList)
		ForestBattleWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, &Monster, 0)
	})
	buttonF.SetBorder(true) //.SetRect(0, 0, 22, 3)

	buttonV := tview.NewButton("go to village").SetSelectedFunc(func() {
		app.Stop()
		Svillage(classes_icons, bg_imgs, player, itemlist, inv)
	})
	buttonV.SetBorder(true) //.SetRect(0, 0, 22, 3)
	// ++++++ flex windows ++++++
	gaucheforet := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gaucheforet.SetBorder(true)
	gaucheforet.AddItem(buttonF, 3, 1, true)
	gaucheforet.AddItem(imageF, 0, 1, false)

	gauchevillage := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gauchevillage.SetBorder(true)
	gauchevillage.AddItem(buttonV, 3, 1, true)
	gauchevillage.AddItem(imageV, 0, 1, false)

	// ++++++ build ++++++
	Gaucheflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headergauche, 3, 1, true).
		AddItem(gauchevillage, 0, 1, true).
		AddItem(gaucheforet, 0, 1, true)

	// =============================== Partie centre ====================================

	//header
	header := tview.NewTextView().SetText("RED PROJECT ULTIMATE")
	header.SetBorder(true)
	header.SetTextColor(redColor)
	header.SetTextAlign(tview.AlignCenter)
	// boite centrale
	centerBox := tview.NewTextView().SetText("")
	centerBox.SetBorder(true)
	centerBox.SetTextAlign(tview.AlignCenter)
	// affichage bas
	quitButton := tview.NewButton("Quitter").
		SetSelectedFunc(func() {
			app.Stop()
		})
	quitButton.SetBorder(true)
	invBoutton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		app.Stop()
		ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
	})
	invBoutton.SetBorder(true)
	statsButton := tview.NewButton("Info Joueur").SetSelectedFunc(func() {
		app.Stop()
		ShowPlayerStats(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
	})
	statsButton.SetBorder(true)
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invBoutton, 4, 2, 1, 2, 0, 0, true).
		AddItem(statsButton, 4, 0, 1, 2, 0, 0, true)
	gridCenter.SetBorder(true)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(centerBox, 0, 1, false).
		AddItem(gridCenter, 0, 1, false)

	// ================================ assemblage ======================================

	//build
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Gaucheflex, 0, 4, true).
		AddItem(Centreflex, 0, 9, true).
		AddItem(Droiteflex, 0, 7, false)

	// =============================== running =============================================
	app.SetFocus(buttonV)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		return 5
		// panic(err4)
	}

	if sceneValue != 0 {
		return sceneValue
	} else {
		return 0
	}
}

/*
Fonction pour afficher l'inventaire.
Génial n'est-ce pas ?
*/
func ShowInventory(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
) {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	consumableButton := tview.NewButton("Consommable").
		SetSelectedFunc(func() {
			app.Stop()
			ShowConsumable(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})

	equipementButton := tview.NewButton("Equipement").
		SetSelectedFunc(func() {
			app.Stop()
			ShowEquipement(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	othersButton := tview.NewButton("Autre").
		SetSelectedFunc(func() {
			app.Stop()
			ShowOthers(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(consumableButton, 4, 3, 1, 1, 0, 0, true).
		AddItem(equipementButton, 4, 2, 1, 1, 0, 0, true).
		AddItem(othersButton, 4, 1, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}

/*
Donne la liste des items consomable
*/
func ShowConsumable(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
) {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	items := tview.NewList()
	CreateItemRow(inv, itemlist, player, true, false, &items, Droiteflex)

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(items, 0, 0, 6, 4, 0, 0, true).
		AddItem(quitButton, 6, 4, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}

/*
Donne la liste des equipements
*/
func ShowEquipement(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
) {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	items := tview.NewList()
	CreateItemRow(inv, itemlist, player, false, true, &items, Droiteflex)

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(items, 0, 0, 6, 4, 0, 0, true).
		AddItem(quitButton, 6, 4, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}

/*
Donne la liste des items non consomable et non equipable
*/
func ShowOthers(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
) int {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// Items de l'inventaire
	items := tview.NewList()
	CreateItemRow(inv, itemlist, player, false, false, &items, Droiteflex)
	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(items, 0, 0, 6, 4, 0, 0, true).
		AddItem(quitButton, 6, 4, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
	return 5
}

// Crée une liste des items, "consumable" doit être mis en true s'il s'agit
// de la page des items consommable, pareil pour "equipable" (les deux ne doivent pas être en 'true')
// et mettre les deux en 'false' si on veut la liste des autres objets.
func CreateItemRow(inv *map[string]int, itemlist map[string]map[string]interface{}, player *map[string]interface{}, consumable bool, equipable bool, list **tview.List, rightPart *tview.Flex) {
	(*list).Clear()
	var Item []string
	var Count []int
	if consumable {
		Item, Count = InventoryTool.GetInventoryConsumables(*inv)
	} else if equipable {
		Item, Count = InventoryTool.GetInventoryEquipable(*inv)
	} else {
		Item, Count = InventoryTool.GetInventoryOther(*inv)
	}
	a := 'a'
	for i := 0; i < len(Item); i++ {
		if Count[i] != 0 {
			var name string
			if !equipable {
				name = itemlist[string(Item[i])]["name"].(string) + ": " + strconv.Itoa(Count[i])
			} else {
				name = itemlist[string(Item[i])]["name"].(string)
			}
			if consumable || equipable {
				(*list).AddItem(name, itemlist[string(Item[i])]["description"].(string), a, func() {
					if consumable {
						combattool.UseConsumable(player, player, Item[i], itemlist, inv)
					} else if equipable {
						InventoryTool.EquipPlayerWith(player, Item[i], inv, itemlist)
					}
					CreateItemRow(inv, itemlist, player, consumable, equipable, list, rightPart)
					updateRightBottomPart(rightPart, *player, itemlist)
				})
			} else {
				(*list).AddItem(name, itemlist[string(Item[i])]["description"].(string), a, nil)
			}
			a += 1

		}
	}
}

func CreateRightPart(
	classes_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
) *tview.Flex {
	// Affichage de l'image
	image := tview.NewImage()
	image.SetImage(classes_icons[(*player)["class"].(string)])
	image.SetBorder(true)

	createRightBottomPart(player, itemlist)

	Droiteflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 0, 1, false).
		AddItem(createRightBottomPart(player, itemlist), 0, 1, false)
	return Droiteflex
}

func createRightBottomPart(
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
) *tview.Grid {

	grid := tview.NewGrid().
		SetRows(3, 0, 0, 0, 0, 3, 0, 3, 0).
		SetColumns(0, 0)
	// Affichage menu equipements et barres de vie et de mana
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text).
			SetDynamicColors(true)
	}

	newBarPrimitive := func(value int, max_val int, colorA string, colorB string) tview.Primitive {
		bar := "╔───────────────────────── / ▲ " + string(byte(92)) + " ─────────────────────────╗\n[white]" + colorA
		tmp := int((float64(value) / float64(max_val)) * float64(55))
		for i := 0; i < tmp; i++ {
			bar += "▮"
		}
		bar += colorB
		for i := tmp; i < 55; i++ {
			bar += "▯"
		}
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(bar + "[white]\n╚───────────────────────── " + string(byte(92)) + " ▼ / ─────────────────────────╝").
			SetDynamicColors(true)
	}

	newPrimitiveEquipmentSlot := func(text string, slot string) tview.Primitive {
		res := text
		if (*player)[slot].(string) != "" {
			res += " " + itemlist[(*player)[slot].(string)]["name"].(string) + "\n"
			switch itemlist[(*player)[slot].(string)]["type"].(string) {
			case "maxhealth":
				res += "+ " + strconv.Itoa(itemlist[(*player)[slot].(string)]["value"].(int)) + "[red]♥"
			case "maxmana":
				res += "+ " + strconv.Itoa(itemlist[(*player)[slot].(string)]["value"].(int)) + "[blue]✦"
			}

		} else {
			res += " Ø\n "
		}
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(res).
			SetDynamicColors(true)
	}

	casque := newPrimitiveEquipmentSlot("Casque :", "EquipmentHead")
	plastron := newPrimitiveEquipmentSlot("Plastron :", "EquipmentTorso")
	jambières := newPrimitiveEquipmentSlot("Jambières :", "EquipmentLegs")
	bottes := newPrimitiveEquipmentSlot("Bottes :", "EquipmentBoots")

	grid.Clear()
	grid.SetTitle("<[ " + (*player)["name"].(string) + " ]>")
	grid.SetBorder(true)
	grid.AddItem(newPrimitive("Equipement :"), 0, 0, 1, 2, 0, 0, false)
	grid.AddItem(newPrimitive("\nPoints de vie : "+strconv.Itoa((*player)["hp"].(int))+" / "+strconv.Itoa((*player)["max_hp"].(int))+"[red] ♥"), 5, 0, 1, 2, 0, 0, false)
	grid.AddItem(newBarPrimitive((*player)["hp"].(int), (*player)["max_hp"].(int), "[green]", "[red]"), 6, 0, 1, 2, 0, 0, false)
	grid.AddItem(newPrimitive("\nMana : "+strconv.Itoa((*player)["mana"].(int))+" / "+strconv.Itoa((*player)["max_mana"].(int))+"[blue] ✦"), 7, 0, 1, 2, 0, 0, false)
	grid.AddItem(newBarPrimitive((*player)["mana"].(int), (*player)["max_mana"].(int), "[blue]", "[gray]"), 8, 0, 1, 2, 0, 0, false)
	grid.AddItem(casque, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(plastron, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(jambières, 3, 0, 1, 1, 0, 0, false)
	grid.AddItem(bottes, 4, 0, 1, 1, 0, 0, false)

	return grid
}

func updateRightBottomPart(
	rightPart *tview.Flex,
	player map[string]interface{},
	itemlist map[string]map[string]interface{},
) {
	rightPart.RemoveItem(rightPart.GetItem(1))
	rightPart.AddItem(createRightBottomPart(&player, itemlist), 0, 1, false)
}

func ShowPlayerStats(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
) {
	app := tview.NewApplication()
	// Partie droite
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	// Partie gauche
	recap := "Nom : " + (*player)["name"].(string)
	recap += "\nClasse : " + classList[(*player)["class"].(string)]["class"].(string)
	recap += "\nPV : " + strconv.Itoa((*player)["hp"].(int)) + " / " + strconv.Itoa((*player)["max_hp"].(int)) + " [red]♥[white]"
	recap += "\nMana : " + strconv.Itoa((*player)["mana"].(int)) + " / " + strconv.Itoa((*player)["max_mana"].(int)) + " [blue]✦[white]"
	recap += "\nOr : " + strconv.Itoa((*player)["gold"].(int))
	recap += "\nSkills : "
	for _, v := range (*player)["skills"].([]string) {
		recap += "\n\u0009- [yellow]" + skillList[v]["name"].(string) + " :[white] "
		switch skillList[v]["type"].(string) {
		case "dmg":
			recap += "Inflige " + strconv.Itoa(skillList[v]["atk_points"].(int)) + " dégats "

		case "heal":
			recap += "Soigne " + strconv.Itoa(skillList[v]["atk_points"].(int)) + " hp "
		}
		if skillList[v]["target_player"].(bool) {
			recap += "au joueur."
		} else {
			recap += "à l'ennemi."
		}
		if skillList[v]["mana_cost"].(int) > 0 {
			recap += " Coute " + strconv.Itoa(skillList[v]["mana_cost"].(int)) + "mana."
		}
	}
	tmp1 := (*player)["max_hp"].(int) - classList[(*player)["class"].(string)]["max_hp"].(int)
	tmp2 := (*player)["max_mana"].(int) - classList[(*player)["class"].(string)]["max_mana"].(int)
	recap += "\nBuffs d'équipement :"
	if tmp1 == 0 && tmp2 == 0 {
		recap += " Aucun " + strconv.Itoa((*player)["max_hp"].(int)) + " " + strconv.Itoa((*player)["max_mana"].(int))
	}
	if tmp1 > 0 {
		recap += "\n\u0009- +" + strconv.Itoa(tmp1) + " [red]♥[white]"
	}
	if tmp2 > 0 {
		recap += "\n\u0009- +" + strconv.Itoa(tmp2) + " [blue]✦[white]"
	}
	playerDesc := tview.NewTextView().SetText(recap)
	playerDesc.SetDynamicColors(true)
	playerDesc.SetBorderColor(tcell.ColorLightYellow)

	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(playerDesc, 0, 0, 4, 5, 0, 0, true)
	gridCenter.SetBorder(true).SetBackgroundColor(tcell.ColorBlack)

	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}

func ForestBattleWindow(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	monster *map[string]interface{},
	turn int,
) {
	buttonActivated := 0
	app := tview.NewApplication()

	// Fonction d'arrêt différé
	differedStop := func(s int) {
		time.Sleep(time.Duration(s) * time.Second)
		app.Stop()
	}

	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	// Boutons
	AttackButton := tview.NewButton("Attaquer")
	BackpackButton := tview.NewButton("Sac à dos")
	FleeButton := tview.NewButton("Prendre la fuite")

	// Monstre Elements + chatbox
	ChatBox := tview.NewTextView()
	MonsterHPBar := tview.NewTextView()
	MonsterHPBar.SetDynamicColors(true).SetTextAlign(tview.AlignCenter)
	MonsterIcon := tview.NewImage()

	// fenêtre Backpack et Attaque
	BackpackMenu := tview.NewList()
	BackpackMenu.SetBorder(true).SetTitle("<[ Sac à dos ]>")
	EmptyBackpack := tview.NewTextView().SetText("\n\n\n\n\nSac à dos vide...")
	EmptyBackpack.SetTextAlign(tview.AlignCenter).SetBorder(true).SetTitle("<[ Sac à dos ]>")
	AttackMenu := tview.NewList()
	AttackMenu.SetBorder(true).SetTitle("<[ Vos skills ]>")

	ChatBox.SetBorder(true).SetTitle("<[ Tour n°" + strconv.Itoa(turn) + ": À vous de jouer ]>")
	if turn == 0 {
		ChatBox.SetText((*monster)["encounter_msg"].(string))
	}
	MonsterHPBar.SetTextAlign(tview.AlignCenter).SetTitle("<[ " + (*monster)["name"].(string) + " ]>").SetBorder(true)
	MonsterIcon.SetImage(monster_icons[(*monster)["id"].(string)])
	MonsterIcon.SetBorder(true)
	// Box Centrale
	gridCenter := tview.NewGrid().
		AddItem(MonsterHPBar, 0, 1, 2, 4, 0, 0, false).
		AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false).
		AddItem(ChatBox, 8, 0, 2, 6, 0, 0, false).
		AddItem(AttackButton, 10, 1, 1, 2, 0, 0, true).
		AddItem(BackpackButton, 10, 3, 1, 2, 0, 0, true).
		AddItem(FleeButton, 11, 2, 1, 2, 0, 0, true)
	gridCenter.SetBorder(true)

	//Fonction du tour du monstre
	monsterTurn := func() {
		time.Sleep(2 * time.Second)
		app.QueueUpdateDraw(func() {
			buttonActivated = 5
			msg, end := combattool.MonsterAttack(player, monster, monsterList, turn)
			updateRightBottomPart(Droiteflex, *player, itemlist)
			if end {
				MonsterIcon.SetImage(bg_imgs["forest"])
				ChatBox.SetText(msg)
				buttonActivated = 6
				go differedStop(5)
			} else if InventoryTool.IsPlayerDead(*player) {
				ChatBox.SetText(msg + "\nVous vous éfondrez sous les coups.")
				buttonActivated = 7
				go differedStop(5)
			} else if combattool.IsMonsterDead(*monster) {
				MonsterIcon.SetImage(bg_imgs["forest"])
				ChatBox.SetText(msg + "\nVotre ennemi tombe au combat.")
				buttonActivated = 8
				go differedStop(5)
			} else {
				ChatBox.SetText(msg)
				go differedStop(2)
			}
		})
	}

	//Fonction refresh bar de vie de l'ennemi
	actuMonsterHPBar := func() {
		bar := "Vie restante : " + strconv.Itoa((*monster)["hp"].(int)) + " / " + strconv.Itoa((*monster)["max_hp"].(int)) + " [red]♥[white]"
		bar += "\n\n╔───────────────────────── / ▲ " + string(byte(92)) + " ─────────────────────────╗\n[white][red]"
		tmp := int((float64((*monster)["hp"].(int)) / float64((*monster)["max_hp"].(int))) * float64(55))
		for i := 0; i < tmp; i++ {
			bar += "▮"
		}
		bar += "[black]"
		for i := tmp; i < 55; i++ {
			bar += "▯"
		}
		MonsterHPBar.Clear()
		MonsterHPBar.SetText(bar + "[white]\n╚───────────────────────── " + string(byte(92)) + " ▼ / ─────────────────────────╝")
	}

	// Donner aux boutons leur fonction

	// Bouton d'attaque
	AttackButton.SetSelectedFunc(func() {
		if buttonActivated == 0 || buttonActivated == 2 || buttonActivated == 3 {
			buttonActivated = 1
			AttackMenu.Clear()
			switch buttonActivated {
			case 0:
				gridCenter.RemoveItem(MonsterIcon)
			case 2:
				gridCenter.RemoveItem(BackpackMenu)
			case 3:
				gridCenter.RemoveItem(EmptyBackpack)
			}
			for i, k := range (*player)["skills"].([]string) {
				name := skillList[k]["name"].(string)
				if skillList[k]["mana_cost"].(int) > 0 {
					name += ". Coup en mana : " + strconv.Itoa(skillList[k]["mana_cost"].(int))
				}
				AttackMenu.AddItem(name, skillList[k]["description"].(string), rune('a'+i), func() {
					if combattool.CanPlayerUseSkill(*player, k, skillList) {
						ChatBox.SetText(skillList[k]["use_text"].(string))
						combattool.UseSkill(player, monster, k, skillList)
						updateRightBottomPart(Droiteflex, *player, itemlist)
						actuMonsterHPBar()
						gridCenter.RemoveItem(AttackMenu)
						gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
						if combattool.IsMonsterDead(*monster) {
							MonsterIcon.SetImage(bg_imgs["forest"])
							ChatBox.SetText("\nVotre ennemi tombe au combat.")
							buttonActivated = 8
							go differedStop(5)
						} else {
							go monsterTurn()
						}
					} else {
						ChatBox.Clear()
						ChatBox.SetText("Vous n'avez pas assez de mana pour lancer : " + skillList[k]["name"].(string))
					}

				})
			}
			gridCenter.AddItem(AttackMenu, 2, 0, 6, 6, 0, 0, true)
			app.SetFocus(AttackMenu)
		}
	})

	// Bouton Sac à dos
	BackpackButton.SetSelectedFunc(func() {
		if buttonActivated >= 0 && buttonActivated <= 1 {
			switch buttonActivated {
			case 0:
				gridCenter.RemoveItem(MonsterIcon)
			case 1:
				gridCenter.RemoveItem(AttackMenu)
			}
			buttonActivated = 2
			ks, vs := InventoryTool.GetInventoryConsumables(*inv)
			if len(ks) == 0 {
				buttonActivated = 3
				gridCenter.AddItem(EmptyBackpack, 2, 0, 6, 6, 0, 0, false)
			} else {
				BackpackMenu.Clear()
				gridCenter.AddItem(BackpackMenu, 2, 0, 6, 6, 0, 0, true)
				for i, k := range ks {
					BackpackMenu.AddItem(itemlist[k]["name"].(string)+" : "+strconv.Itoa(vs[i]), itemlist[k]["description"].(string), rune('a'+i), func() {
						ChatBox.SetText(itemlist[k]["use_text"].(string))
						combattool.UseConsumable(player, monster, k, itemlist, inv)
						updateRightBottomPart(Droiteflex, *player, itemlist)
						actuMonsterHPBar()
						gridCenter.RemoveItem(BackpackMenu)
						gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
						if combattool.IsMonsterDead(*monster) {
							MonsterIcon.SetImage(bg_imgs["forest"])
							ChatBox.SetText("\nVotre ennemi tombe au combat.")
							buttonActivated = 8
							go differedStop(5)
						} else {
							go monsterTurn()
						}
					})
				}
				app.SetFocus(BackpackMenu)
			}
		}
	})

	// Bouton de fuite
	FleeButton.SetSelectedFunc(func() {
		if buttonActivated >= 0 && buttonActivated <= 3 {
			switch buttonActivated {
			case 1:
				{
					gridCenter.RemoveItem(AttackMenu)
					gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
				}
			case 2:
				{
					gridCenter.RemoveItem(BackpackMenu)
					gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
				}
			case 3:
				{
					gridCenter.RemoveItem(EmptyBackpack)
					gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)

				}
			}
			buttonActivated = 4
			ChatBox.SetText("Vous prenez la fuite...")
			go differedStop(3)
		}
	})

	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)

	actuMonsterHPBar()
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
	switch buttonActivated {
	case 4:
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
	case 5:
		ForestBattleWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, monster, turn+1)
	case 6:
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
	case 7:
		GameOverWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, *monster, turn)
	case 8:
		VictoryWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, *monster, turn)
	}
}

func GameOverWindow(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	monster map[string]interface{},
	turn int,
) {
	app := tview.NewApplication()

	restart_bouton := tview.NewButton("Continuer").SetSelectedFunc(func() {
		InventoryTool.HealPlayer(player, (*player)["max_hp"].(int)/2)
		InventoryTool.HealPlayerMana(player, (*player)["ma_mana"].(int))
		InventoryTool.RemoveGoldFromPlayer(player, (*player)["gold"].(int)/3)
		app.Stop()
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
	})
	gameover_msg := "Vous êtes tomber au combat face à ennemi (" + monster["name"].(string) + "). Vous perder " + strconv.Itoa((*player)["gold"].(int)/3) + "or."
	gameover_msg += "\nVous regagnez la moitier de vos points de vie ainsi que votre mana. Faite plus attention à l'avenir..."
	image_gameover := tview.NewImage()
	chatbox := tview.NewTextView().SetText(gameover_msg)
	chatbox.SetDynamicColors(true).SetBorder(true)
	restart_bouton.SetBorder(true)
	image_gameover.SetBorder(true)

	image_gameover.SetImage(bg_imgs["gameover"])

	grid := tview.NewGrid().
		AddItem(image_gameover, 0, 0, 9, 6, 0, 0, false).
		AddItem(chatbox, 9, 0, 1, 5, 0, 0, false).
		AddItem(restart_bouton, 9, 5, 1, 1, 0, 0, true)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func VictoryWindow(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	monster map[string]interface{},
	turn int,
) {
	return
}

func CreateLeftPart(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
	imageTop map[string]image.Image,
	imageBottom map[string]image.Image,
	nameTop string,
	nameBottom string,
	headerText string,
	app *tview.Application,
) (*tview.Flex, *tview.Button, *tview.Button) {

	headergauche := tview.NewTextView().SetText(headerText)
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	buttonTOP := tview.NewButton(nameTop)
	buttonTOP.SetBorder(true)

	buttonBottom := tview.NewButton(nameBottom)
	buttonBottom.SetBorder(true)

	flexTOP := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(buttonTOP, 3, 1, true)
		// futur une imageTop
	flexTOP.SetBorder(true)

	flexBottom := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(buttonBottom, 3, 1, true)
		// futur une imageBottom
	flexBottom.SetBorder(true)

	GaucheFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headergauche, 3, 1, true).
		AddItem(flexTOP, 0, 1, true).
		AddItem(flexBottom, 0, 1, true)

	return GaucheFlex, buttonTOP, buttonBottom
}

func Svillage(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================

	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	//================================PARTIT GAUCHE ===============================

	// linkTop := Sshopvillage(classes_icons, bg_imgs, player, itemlist, inv, "forgeron")
	// linkBottom := Sshopvillage(classes_icons, bg_imgs, player, itemlist, inv, "vendeur")

	Gaucheflex, buttonTOP, buttonBottom := CreateLeftPart(classes_icons, bg_imgs, player, itemlist, inv, bg_imgs, bg_imgs,
		"top", "bottom", "Rendre visite a votre tres cher :", app)

	buttonTOP.SetSelectedFunc(func() {
		app.Stop()
		Sshopvillage(classes_icons, bg_imgs, player, itemlist, inv, "forgeron")
		// if err4 := app.SetRoot(linkTop, true).EnableMouse(true).Run(); err4 != nil {
		// 	panic(err4)
		// }
		//forgeron
	})

	buttonBottom.SetSelectedFunc(func() {
		app.Stop()
		Sshopvillage(classes_icons, bg_imgs, player, itemlist, inv, "vendeur")
		// if err4 := app.SetRoot(linkBottom, true).EnableMouse(true).Run(); err4 != nil {
		// 	panic(err4)
		// }
		//vendeur
	})

	// ============================ CENTRE =============================

	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, player, itemlist, inv)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	invBoutton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		app.Stop()
		ShowInventory(classes_icons, bg_imgs, player, itemlist, inv)
	})
	invBoutton.SetBorder(true)

	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invBoutton, 4, 2, 1, 2, 0, 0, true)

	textTop := tview.NewTextView().SetText("je suis un texte qui raconte une histoire tres tres interessente, telement que vous meme avez oublié pouquoi vous liser ceci")
	textTop.SetBorder(true)
	textTop.SetTextColor(tcell.ColorDarkRed)
	textTop.SetTextAlign(tview.AlignCenter)

	CentreTop := tview.NewFlex().
		AddItem(textTop, 0, 1, false)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(CentreTop, 0, 1, true).
		AddItem(gridCenter, 0, 1, true)

	// ====================== BUILD ==============================
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Gaucheflex, 0, 1, true).
		AddItem(Centreflex, 0, 1, true).
		AddItem(Droiteflex, 0, 1, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}

func Sshopvillage(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
	who string,
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================

	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	//================================PARTIT GAUCHE ===============================

	textTopgauche := tview.NewTextView()
	textTopgauche.SetBorder(true)
	textTopgauche.SetTextColor(tcell.ColorDarkRed)
	textTopgauche.SetTextAlign(tview.AlignCenter)

	switch who { // who is selling
	case "vendeur":
		//image vendeur
		textTopgauche.SetText("hohoho, ici je vend bien de bonne chose, tout cela est t'a toi, mais seulement SI tu ne possede assez de pieces d'or hihihi")

		// liste interactive vendeur
	case "forgeron":
		//image forgerons
		textTopgauche.SetText("Hmfmfhmfm, armes cassé, objet a bricolé, bienvenu dans les flamme et le fer, rien de mieux pour battre l'enfer mfmfmfm")

		// liste interactive for
	}

	Gaucheflex := tview.NewFlex().
		AddItem(textTopgauche, 0, 1, true)
	// ============================ CENTRE =============================
	//bottom
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, player, itemlist, inv)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	invBoutton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		app.Stop()
		ShowInventory(classes_icons, bg_imgs, player, itemlist, inv)
	})
	invBoutton.SetBorder(true)

	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invBoutton, 4, 2, 1, 2, 0, 0, true)
	// top
	// A REMPLAC2 PAR IMAGE DU VILLAGE EN GRAND
	textTop := tview.NewTextView().SetText("je suis un texte qui raconte une histoire tres tres interessente, telement que vous meme avez oublié pouquoi vous liser ceci")
	textTop.SetBorder(true)
	textTop.SetTextColor(tcell.ColorDarkRed)
	textTop.SetTextAlign(tview.AlignCenter)

	CentreTop := tview.NewFlex().
		AddItem(textTop, 0, 1, false)
	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(CentreTop, 0, 1, true).
		AddItem(gridCenter, 0, 1, true)

	// ============================ build ===========================================
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Gaucheflex, 0, 1, true).
		AddItem(Centreflex, 0, 1, true).
		AddItem(Droiteflex, 0, 1, false)

	// return Machted
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}

}
