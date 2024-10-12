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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
) int {
	sceneValue := 0

	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	// =============================== Partie gauche ====================================

	// ++++++ header ++++++
	headergauche := tview.NewTextView().SetText("Intitié le voyage en direction :")
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	// ++++++ images ++++++
	imageV := tview.NewImage()
	imageV.SetImage(bg_imgs["village_short"])

	imageF := tview.NewImage()
	imageF.SetImage(bg_imgs["forest_short"])

	// ++++++ Boutons ++++++
	buttonF := tview.NewButton("De la Sombre fôret").SetSelectedFunc(func() {
		sceneValue = 2
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		Monster := combattool.GenRandMonster(monsterList)
		ForestBattleWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, &Monster, craftList, tradeList, donjonLevels)
	})
	buttonF.SetBorder(true) //.SetRect(0, 0, 22, 3)

	buttonV := tview.NewButton("Du grand village").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		Svillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
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
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
		})
	quitButton.SetBorder(true)
	invBoutton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	})
	invBoutton.SetBorder(true)
	statsButton := tview.NewButton("Info Joueur").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		ShowPlayerStats(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	})
	statsButton.SetBorder(true)

	// testDonjon := tview.NewButton("Test Donjon").SetSelectedFunc(func() {
	// 	InventoryTool.PlaySound("ressource/sound_button.mp3")
	// 	app.Stop()
	// 	EnterDonjon(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, "sky_tower")
	// })
	// statsButton.SetBorder(true)

	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invBoutton, 4, 2, 1, 2, 0, 0, true).
		AddItem(statsButton, 4, 0, 1, 2, 0, 0, true) //.
		//AddItem(testDonjon, 2, 0, 1, 2, 0, 0, true)
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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	sourcepage func(),
) {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			sourcepage()
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	consumableButton := tview.NewButton("Consommable").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowConsumable(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, sourcepage)
		})

	equipementButton := tview.NewButton("Equipement").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowEquipement(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, sourcepage)
		})
	othersButton := tview.NewButton("Autre").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowOthers(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, sourcepage)
		})

	// image inventaire

	imageInventaire := tview.NewImage()
	imageInventaire.SetImage(bg_imgs["inventory"])

	// Boite text centrale
	barCentrale := tview.NewTextView().SetText("[yellow]Vous jetez un œuil dans votre sac. Que cherchez vous ?[white]")
	barCentrale.SetTextAlign(tview.AlignCenter).SetDynamicColors(true).SetBorder(true)

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(imageInventaire, 0, 0, 5, 4, 0, 0, true).
		AddItem(barCentrale, 5, 0, 1, 4, 0, 0, true).
		AddItem(quitButton, 6, 3, 1, 1, 0, 0, true).
		AddItem(consumableButton, 6, 2, 1, 1, 0, 0, true).
		AddItem(equipementButton, 6, 1, 1, 1, 0, 0, true).
		AddItem(othersButton, 6, 0, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true).
		SetTitle(" < Inventaire > ")
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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	sourcepage func(),
) {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, sourcepage)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	items := tview.NewList()
	CreateItemRow(inv, itemlist, player, true, false, &items, Droiteflex)

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(items, 0, 0, 4, 4, 0, 0, true).
		AddItem(quitButton, 4, 3, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true).
		SetTitle(" < Consommable > ")
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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	sourcepage func(),
) {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, sourcepage)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	items := tview.NewList()
	CreateItemRow(inv, itemlist, player, false, true, &items, Droiteflex)

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(items, 0, 0, 4, 4, 0, 0, true).
		AddItem(quitButton, 4, 3, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true).
		SetTitle(" < Equipement > ")
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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	sourcepage func(),
) int {
	app := tview.NewApplication()
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, sourcepage)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// Items de l'inventaire
	items := tview.NewList()
	CreateItemRow(inv, itemlist, player, false, false, &items, Droiteflex)
	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(items, 0, 0, 4, 4, 0, 0, true).
		AddItem(quitButton, 4, 3, 1, 1, 0, 0, true)
	gridCenter.SetBorder(true).
		SetTitle(" < Autre > ")
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
	weapon := newPrimitiveEquipmentSlot("Arme :", "EquipmentWeapon")

	gold := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("OR : " + strconv.Itoa((*player)["gold"].(int)) + " [yellow]⬤[white]").
		SetDynamicColors(true)

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
	grid.AddItem(weapon, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(gold, 4, 1, 1, 1, 0, 0, false)

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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	sourcepage func(),
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
		lst_atk_points := skillList[v]["atk_points"].([]interface{})
		lst_targets := skillList[v]["target_player"].([]interface{})
		if skillList[v]["mana_cost"].(int) > 0 {
			recap += " Coute " + strconv.Itoa(skillList[v]["mana_cost"].(int)) + " mana."
		}
		for i, str := range skillList[v]["type"].([]interface{}) {
			switch str.(string) {
			case "dmg":
				recap += "\n\u0009\u0009- Inflige " + strconv.Itoa(int(lst_atk_points[i].(float64))) + " dégats "

			case "heal":
				recap += "\n\u0009\u0009- Soigne " + strconv.Itoa(int(lst_atk_points[i].(float64))) + " hp "
			}
			if lst_targets[i].(bool) {
				recap += "au joueur."
			} else {
				recap += "à l'ennemi."
			}
		}
	}
	tmp1 := (*player)["max_hp"].(int) - classList[(*player)["class"].(string)]["max_hp"].(int)
	tmp2 := (*player)["max_mana"].(int) - classList[(*player)["class"].(string)]["max_mana"].(int)
	recap += "\nBuffs d'équipement :"
	if tmp1 == 0 && tmp2 == 0 {
		recap += " Aucun"
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
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			sourcepage()
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
	lootList map[string]map[string]map[string][]int,
	monster *map[string]interface{},
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
) {
	var turn int = 0
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
	ChatBox.SetBorder(true).SetTitle("<[ Tour n°" + strconv.Itoa(turn) + " : À vous de jouer ]>")
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
	var actuMonsterHPBarRecur func(int)
	actuMonsterHPBarRecur = func(n int) {
		if n > 0 {
			actuMonsterHPBar()
			time.Sleep(1 * time.Second)
			go actuMonsterHPBarRecur(n - 1)
		}
	}

	//Fonction de reset de la page pour un nouveau tour
	reset := func() {
		time.Sleep(3 * time.Second)
		actuMonsterHPBar()
		turn++
		ChatBox.SetBorder(true).SetTitle("<[ Tour n°" + strconv.Itoa(turn) + " : À vous de jouer ]>")
		ChatBox.SetText("Que faite vous ?")
		buttonActivated = 0
	}

	//Fonction du tour du monstre
	monsterTurn := func() {
		time.Sleep(3 * time.Second)
		app.QueueUpdateDraw(func() {
			buttonActivated = 5
			actuMonsterHPBar()
			msg, end := combattool.MonsterAttack(player, monster, monsterList, turn)
			updateRightBottomPart(Droiteflex, *player, itemlist)
			ChatBox.SetBorder(true).SetTitle("<[ Tour n°" + strconv.Itoa(turn) + " : Au tour de l'ennemi ]>")
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
				go reset()
			}
		})
	}

	// Donner aux boutons leur fonction

	// Bouton d'attaque
	AttackButton.SetSelectedFunc(func() {
		if buttonActivated == 0 || buttonActivated == 2 || buttonActivated == 3 {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
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
						buttonActivated = 5
						tmp := ""
						tmp += skillList[k]["use_text"].(string)
						lst_atk_points := skillList[k]["atk_points"].([]interface{})
						for i, k := range skillList[k]["type"].([]interface{}) {
							switch k.(string) {
							case "dmg":
								{
									tmp += " Vous infliguez "
									if (*monster)["spe"].(string) == "reduce_dmg" {
										tmp += strconv.Itoa(int(lst_atk_points[i].(float64))-(*monster)["special"].(int)) + " dégats." + "\nL'ennemi à une peau renforcé, ces dégats subies sont réduits.\n"
									} else {
										tmp += strconv.Itoa(int(lst_atk_points[i].(float64))) + " dégats.\n"
									}
								}
							case "heal":
								{
									tmp += "Vous récuperez " + strconv.Itoa(int(lst_atk_points[i].(float64))) + " points de vie.\n"
								}
							}
						}
						ChatBox.SetText(tmp)
						combattool.UseSkill(player, monster, k, skillList)
						updateRightBottomPart(Droiteflex, *player, itemlist)
						actuMonsterHPBar()
						gridCenter.RemoveItem(AttackMenu)
						gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
						if InventoryTool.IsPlayerDead(*player) {
							end := func() {
								app.QueueUpdateDraw(func() {
									time.Sleep(3 * time.Second)
									ChatBox.SetText("\nVous vous effondrez sous les coups.")
									buttonActivated = 7
									go differedStop(5)
								})
							}
							go end()
						} else if combattool.IsMonsterDead(*monster) {
							end := func() {
								app.QueueUpdateDraw(func() {
									time.Sleep(3 * time.Second)
									MonsterIcon.SetImage(bg_imgs["forest"])
									ChatBox.SetText("\nVotre ennemi tombe au combat.")
									buttonActivated = 8
									go differedStop(5)
								})
							}
							go end()
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
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			switch buttonActivated {
			case 0:
				gridCenter.RemoveItem(MonsterIcon)
			case 1:
				gridCenter.RemoveItem(AttackMenu)
			}
			buttonActivated = 2
			ks, vs := InventoryTool.GetInventoryConsumablesForCombat(*inv)
			if len(ks) == 0 {
				buttonActivated = 3
				gridCenter.AddItem(EmptyBackpack, 2, 0, 6, 6, 0, 0, false)
			} else {
				BackpackMenu.Clear()
				gridCenter.AddItem(BackpackMenu, 2, 0, 6, 6, 0, 0, true)
				for i, k := range ks {
					BackpackMenu.AddItem(itemlist[k]["name"].(string)+" : "+strconv.Itoa(vs[i]), itemlist[k]["description"].(string), rune('a'+i), func() {
						buttonActivated = 5
						ChatBox.SetText(itemlist[k]["use_text"].(string))
						combattool.UseConsumable(player, monster, k, itemlist, inv)
						updateRightBottomPart(Droiteflex, *player, itemlist)
						actuMonsterHPBar()
						gridCenter.RemoveItem(BackpackMenu)
						gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
						if InventoryTool.IsPlayerDead(*player) {
							end := func() {
								app.QueueUpdateDraw(func() {
									time.Sleep(3 * time.Second)
									ChatBox.SetText("\nVous vous éfondrez sous les coups.")
									buttonActivated = 7
									go differedStop(5)
								})
							}
							go end()
						} else if combattool.IsMonsterDead(*monster) {
							end := func() {
								app.QueueUpdateDraw(func() {
									time.Sleep(3 * time.Second)
									MonsterIcon.SetImage(bg_imgs["forest"])
									ChatBox.SetText("\nVotre ennemi tombe au combat.")
									buttonActivated = 8
									go differedStop(5)
								})
							}
							go end()
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
			InventoryTool.PlaySound("ressource/sound_button.mp3")
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
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
	case 6:
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
	case 7:
		GameOverWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, *monster, turn)
	case 8:
		VictoryWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, *monster, turn)
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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	monster map[string]interface{},
	turn int,
) {
	app := tview.NewApplication()
	InventoryTool.HealPlayer(player, (*player)["max_hp"].(int)/2)
	InventoryTool.HealPlayerMana(player, (*player)["max_mana"].(int))
	InventoryTool.RemoveGoldFromPlayer(player, (*player)["gold"].(int)/3)

	restart_bouton := tview.NewButton("Continuer").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
	})
	gameover_msg := "Vous êtes tombé(e) au combat face à ennemi (" + monster["name"].(string) + "). Vous perdez " + strconv.Itoa((*player)["gold"].(int)/3) + " or."
	gameover_msg += "\nVous regagnez la moitié de vos points de vie ainsi que votre mana. Faites plus attention à l'avenir..."
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
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	monster map[string]interface{},
	turn int,
) {
	app := tview.NewApplication()
	loot := combattool.GetCombatLoot(monster["id"].(string), inv, player, lootList, itemlist)
	InventoryTool.HealPlayerMana(player, (*player)["max_mana"].(int))

	restart_bouton := tview.NewButton("Continuer").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
	})
	gameover_msg := "Vous avez vaillamment vaincu l'ennemi (" + monster["name"].(string) + ")."
	if loot != "" {
		gameover_msg += "\nVous trouvez " + loot + "."
	}
	gameover_msg += "\nVous sortez victorieux de la forêt, un sourire en coin."
	image_gameover := tview.NewImage()
	chatbox := tview.NewTextView().SetText(gameover_msg)
	chatbox.SetDynamicColors(true).SetBorder(true)
	restart_bouton.SetBorder(true)
	image_gameover.SetBorder(true)

	image_gameover.SetImage(bg_imgs["forest"])

	grid := tview.NewGrid().
		AddItem(image_gameover, 0, 0, 9, 6, 0, 0, false).
		AddItem(chatbox, 9, 0, 1, 5, 0, 0, false).
		AddItem(restart_bouton, 9, 5, 1, 1, 0, 0, true)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		print("\n\nICI!\n\n")
		panic(err)
	}
}

func CreateLeftPart(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
	imageTop string,
	imageBottom string,
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

	imageT := tview.NewImage()
	imageT.SetImage(bg_imgs[imageTop])
	imageB := tview.NewImage()
	imageB.SetImage(bg_imgs[imageBottom])

	flexTOP := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(imageT, 0, 1, true).
		AddItem(buttonTOP, 3, 1, true)
	flexTOP.SetBorder(true)

	flexBottom := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(imageB, 0, 1, true).
		AddItem(buttonBottom, 3, 1, true)
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
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================

	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	//================================PARTIT GAUCHE ===============================

	Gaucheflex, buttonTOP, buttonBottom := CreateLeftPart(classes_icons, bg_imgs, player, itemlist, *inv, "forge", "merchant",
		"Le Forgeron", "Le Marchant", "Rendre visite à votre connaisance :", app)

	buttonTOP.SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		Sshopvillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, "forgeron")
	})

	buttonBottom.SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		Sshopvillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, "vendeur_buy")
	})

	// ============================ CENTRE =============================

	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	invBoutton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
			Svillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	})
	invBoutton.SetBorder(true)

	statsButton := tview.NewButton("Stats").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		ShowPlayerStats(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
			Svillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	})
	invBoutton.SetBorder(true)

	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invBoutton, 4, 2, 1, 2, 0, 0, true).
		AddItem(statsButton, 4, 0, 1, 2, 0, 0, true)

	imageGauche := tview.NewImage()
	imageGauche.SetBorder(true)
	imageGauche.SetImage(bg_imgs["village"])

	CentreTop := tview.NewFlex().
		// AddItem(textTop, 0, 1, false)
		AddItem(imageGauche, 0, 1, true)

	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(CentreTop, 0, 1, true).
		AddItem(gridCenter, 0, 1, true)

	// ====================== BUILD ==============================
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Gaucheflex, 45, 1, true).
		AddItem(Centreflex, 0, 1, true).
		AddItem(Droiteflex, 0, 1, false)
	if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}

func Sshopvillage(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	who string,
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================

	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	//================================PARTIT GAUCHE ===============================

	imageBottom := tview.NewImage()
	imageBottom.SetBorder(true)

	shop := tview.NewList()
	shop.SetBorder(true)
	//CHOIX D IMAGE FAIT DANS SWITCH CENTRE

	Gaucheflex := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0)
	Gaucheflex.
		AddItem(shop, 2, 0, 2, 4, 0, 0, true).
		AddItem(imageBottom, 0, 0, 2, 4, 0, 0, true)

	// ============================ CENTRE =============================
	//bottom
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			Svillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	invButton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
			Sshopvillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, who)
		})
	})
	invButton.SetBorder(true)

	statsButton := tview.NewButton("Stats").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		ShowPlayerStats(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
			Sshopvillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, who)
		})
	})
	statsButton.SetBorder(true)

	sellButton := tview.NewButton("Vendre").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		Sshopvillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, "vendeur_sell")
	})

	buyButton := tview.NewButton("Acheter").SetSelectedFunc(func() {
		InventoryTool.PlaySound("ressource/sound_button.mp3")
		app.Stop()
		Sshopvillage(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, "vendeur_buy")
	})

	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invButton, 4, 2, 1, 2, 0, 0, true).
		AddItem(statsButton, 4, 0, 1, 2, 0, 0, true)
	// top

	textTop := tview.NewTextView()
	textTop.SetBorder(true)
	textTop.SetTextColor(tcell.ColorDarkRed)
	textTop.SetTextAlign(tview.AlignCenter)

	switch who { // who is selling + IMAGE GAUCHE BOTTOM
	case "vendeur_buy":
		{
			textTop.SetText("Hohoho, ici je vend bien de bonnes choses, tout cela est à toi, mais seulement SI tu possède assez de pièces d'or hihihi...")
			imageBottom.SetImage(bg_imgs["merchant"])
			ShowShopBuy(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, &shop, Gaucheflex, gridCenter, Droiteflex)
			gridCenter.AddItem(sellButton, 0, 4, 4, 1, 0, 0, true)
		}

	case "vendeur_sell":
		{
			textTop.SetText("Si tu as des objets de valeurs, je veut bien y jeter un coup d'oeuil, sinon vat-en !")
			imageBottom.SetImage(bg_imgs["merchant"])
			ShowShopSell(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, &shop, Gaucheflex, gridCenter, Droiteflex)
			gridCenter.AddItem(buyButton, 0, 4, 4, 1, 0, 0, true)
		}

	case "forgeron":
		{
			textTop.SetText("Hpmf hpmf hpmf, armes cassées? objets à bricoler? Bienvenue dans les flammes et le fer, rien de mieux pour battre l'enfer pfmh pfmh pfmh...")
			imageBottom.SetImage(bg_imgs["forge"])
			ShowForge(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, &shop, Gaucheflex, gridCenter, Droiteflex)
		}
	}

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
	tradeList map[string]map[string]map[string]int,
	list **tview.List,
	bottomGrid *tview.Grid,
	midGrid *tview.Grid,
	bottomRightPart *tview.Flex,
) {
	(*list).Clear()
	Craftable := []string{}
	for k := range craftList {
		Craftable = append(Craftable, k)
	}
	a := 'a'
	I := false
	for _, i := range Craftable {
		name := itemlist[string(i)]["name"].(string)
		if !InventoryTool.CanCraft(i, craftList, *inv, *player) {
			name += " ❌"
			(*list).AddItem(name, itemlist[string(i)]["description"].(string), a, nil)
		} else {
			(*list).AddItem(name, itemlist[string(i)]["description"].(string), a, func() {
				InventoryTool.Craft(i, craftList, inv, player)
				UpdateBottomGrid(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, list, bottomGrid, midGrid, 0, bottomRightPart)
			})
		}
		a++
		if !I {
			I = true
			ingredient := []string{}
			for k := range craftList[i] {
				ingredient = append(ingredient, k)
			}
			recipe_text := ""
			for _, k := range ingredient {
				if k != "cvalue" {
					recipe_text += "  - " + strconv.Itoa(craftList[i][k]) + " " + itemlist[k]["name"].(string) + "\n"
				}
			}
			text := "\n  Coût:\n" + "  - " + strconv.Itoa(craftList[i]["cvalue"]) + "💰\n" + recipe_text
			midGrid.AddItem(tview.NewTextView().SetText(text), 0, 0, 4, 5, 0, 0, true)
		}
	}
	(*list).SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		itemId := Craftable[index]
		ingredient := []string{}
		for k := range craftList[itemId] {
			ingredient = append(ingredient, k)
		}
		recipe_text := ""
		for _, k := range ingredient {
			if k != "cvalue" {
				recipe_text += "  - " + strconv.Itoa(craftList[itemId][k]) + " " + itemlist[k]["name"].(string) + "\n"
			}
		}
		text := "\n  Coût:\n" + "  - " + strconv.Itoa(craftList[itemId]["cvalue"]) + "💰\n" + recipe_text
		midGrid.AddItem(tview.NewTextView().SetText(text), 0, 0, 4, 5, 0, 0, true)
	})
}

func UpdateBottomGrid(
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
	tradeList map[string]map[string]map[string]int,
	list **tview.List,
	bottomGrid *tview.Grid,
	midGrid *tview.Grid,
	shop_type int,
	bottomRightPart *tview.Flex,

) {
	bottomGrid.RemoveItem(*list)
	updateRightBottomPart(bottomRightPart, *player, itemlist)
	switch shop_type {
	case 0:
		ShowForge(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, list, bottomGrid, midGrid, bottomRightPart)
	case 1:
		ShowShopBuy(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, list, bottomGrid, midGrid, bottomRightPart)
	case 2:
		ShowShopSell(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, list, bottomGrid, midGrid, bottomRightPart)
	}

	bottomGrid.AddItem(*list, 2, 0, 2, 4, 0, 0, true)
}

func ShowShopBuy(
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
	tradeList map[string]map[string]map[string]int,
	list **tview.List,
	bottomGrid *tview.Grid,
	midGrid *tview.Grid,
	bottomRightPart *tview.Flex,
) {
	(*list).Clear()
	a := 'a'
	Items := []string{}
	I := false
	for i := range tradeList["merchant"]["sells"] {
		Items = append(Items, i)
		name := itemlist[i]["name"].(string)
		if !InventoryTool.CanBuyItemFrom(*player, tradeList, "merchant", i) {
			name += " ❌"
			(*list).AddItem(name, itemlist[i]["description"].(string), a, nil)
		} else {
			(*list).AddItem(name, itemlist[i]["description"].(string), a, func() {
				InventoryTool.BuyItemFrom(player, inv, tradeList, "merchant", i)
				if i == "CP_Heal" && tradeList["merchant"]["sells"]["CP_Heal"] == 0 {
					tradeList["merchant"]["sells"]["CP_Heal"] = 3
				}
				UpdateBottomGrid(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, list, bottomGrid, midGrid, 1, bottomRightPart)
			})
		}
		a++
		if !I {
			I = true
			text := "\n  Prix Rachat :\n" + "  - +" + strconv.Itoa(tradeList["merchant"]["buys"][i]) + "💰\n"
			val, ok := (*inv)[i]
			if ok {
				text += "\n\n Dans le sac : " + strconv.Itoa(val)
			} else {
				text += "\n\n Dans le sac : 0"
			}
			midGrid.AddItem(tview.NewTextView().SetText(text), 0, 0, 4, 4, 0, 0, true)
		}
	}
	(*list).SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		itemId := Items[index]
		text := "\n  Coût :\n" + "  - " + strconv.Itoa(tradeList["merchant"]["sells"][itemId]) + "💰\n"
		val, ok := (*inv)[itemId]
		if ok {
			text += "\n\n Dans le sac : " + strconv.Itoa(val)
		} else {
			text += "\n\n Dans le sac : 0"
		}
		midGrid.AddItem(tview.NewTextView().SetText(text), 0, 0, 4, 4, 0, 0, true)
	})

}

func ShowShopSell(
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
	tradeList map[string]map[string]map[string]int,
	list **tview.List,
	bottomGrid *tview.Grid,
	midGrid *tview.Grid,
	bottomRightPart *tview.Flex,
) {
	(*list).Clear()
	a := 'a'
	Items := []string{}
	I := false
	for i := range tradeList["merchant"]["buys"] {
		Items = append(Items, i)
		name := itemlist[i]["name"].(string)
		if !InventoryTool.CanSellItemTo(*inv, tradeList, "merchant", i) {
			name += " ❌"
			(*list).AddItem(name, itemlist[i]["description"].(string), a, nil)
		} else {
			(*list).AddItem(name, itemlist[i]["description"].(string), a, func() {
				InventoryTool.SellItemTo(player, inv, tradeList, "merchant", i)
				UpdateBottomGrid(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, craftList, tradeList, list, bottomGrid, midGrid, 2, bottomRightPart)
			})
		}
		a++
		if !I {
			I = true
			text := "\n  Prix Rachat :\n" + "  - +" + strconv.Itoa(tradeList["merchant"]["buys"][i]) + "💰\n"
			val, ok := (*inv)[i]
			if ok {
				text += "\n\n Dans le sac : " + strconv.Itoa(val)
			} else {
				text += "\n\n Dans le sac : 0"
			}
			midGrid.AddItem(tview.NewTextView().SetText(text), 0, 0, 4, 4, 0, 0, true)
		}
	}
	(*list).SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		itemId := Items[index]
		text := "\n  Prix Rachat :\n" + "  - +" + strconv.Itoa(tradeList["merchant"]["buys"][itemId]) + "💰\n"
		val, ok := (*inv)[itemId]
		if ok {
			text += "\n\n Dans le sac : " + strconv.Itoa(val)
		} else {
			text += "\n\n Dans le sac : 0"
		}
		midGrid.AddItem(tview.NewTextView().SetText(text), 0, 0, 4, 4, 0, 0, true)
	})
}

func EnterDonjon(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	monster_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv *map[string]int,
	classList map[string]map[string]interface{},
	skillList map[string]map[string]interface{},
	monsterList map[string]map[string]interface{},
	lootList map[string]map[string]map[string][]int,
	craftList map[string]map[string]int,
	tradeList map[string]map[string]map[string]int,
	donjonLevels map[string]map[string]map[string]interface{},
	donjonID string,
) {
	app := tview.NewApplication()

	// ===================================== Partie droite =====================================

	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	// ===================================== Partie gauche =====================================

	exitDonjonButton := tview.NewButton("Sortir de la tour").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	exitDonjonButton.SetBackgroundColor(tcell.ColorOrange)
	exitDonjonButton.SetBorderColor(tcell.ColorOrange)
	exitDonjonButton.SetBorder(true)

	playerStatButton := tview.NewButton("Stats").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels)
		})
	playerStatButton.SetBorder(true)

	playerInventoryButton := tview.NewButton("Inventaire").
		SetSelectedFunc(func() {
			InventoryTool.PlaySound("ressource/sound_button.mp3")
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, func() {
				EnterDonjon(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, lootList, craftList, tradeList, donjonLevels, donjonID)
			})
		})
	playerInventoryButton.SetBorder(true)

	gridCenter := tview.NewGrid().
		AddItem(exitDonjonButton, 11, 4, 1, 2, 0, 0, true).
		AddItem(playerInventoryButton, 11, 4, 1, 2, 0, 0, true).
		AddItem(playerStatButton, 11, 0, 1, 2, 0, 0, true)

	// ========================================= Build =========================================

	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gridCenter, 0, 1, true)
	Centreflex.SetBorder(true)

	build := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Centreflex, 0, 13, true).
		AddItem(Droiteflex, 0, 7, false)

	if err4 := app.SetRoot(build, true).EnableMouse(true).Run(); err4 != nil {
		panic(err4)
	}
}
