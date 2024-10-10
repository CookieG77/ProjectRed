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
	running := true

	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)

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
		running = false
		Monster := combattool.GenRandMonster(monsterList)
		ForestBattleWindow(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList, &Monster, 0)
	})
	buttonF.SetBorder(true) //.SetRect(0, 0, 22, 3)

	buttonV := tview.NewButton("go to village").SetSelectedFunc(func() {
		sceneValue = 1
		app.Stop()
		running = false
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
			running = false
		})
	quitButton.SetBorder(true)
	invBoutton := tview.NewButton("Inventaire").SetSelectedFunc(func() {
		app.Stop()
		running = false
		ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
	})
	invBoutton.SetBorder(true)
	statsButton := tview.NewButton("Info Joueur").SetSelectedFunc(func() {
		app.Stop()
		running = false
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
	running := true
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
			SmenuRender(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	consumableButton := tview.NewButton("Consommable").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
			ShowConsumable(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})

	equipementButton := tview.NewButton("Equipement").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
			ShowEquipement(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	othersButton := tview.NewButton("Autre").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
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
	running := true
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	app.Stop()
	running = false

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true)
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
	running := true
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	app.Stop()
	running = false

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true)
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
	running := true
	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)
	// Boutons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			running = false
			ShowInventory(classes_icons, bg_imgs, monster_icons, player, itemlist, inv, classList, skillList, monsterList)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	app.Stop()
	running = false

	// Box Centrale
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true)
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

func CreateRightPart(
	classes_icons map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	running *bool,
	app *tview.Application,
) *tview.Flex {
	// Affichage de l'image
	image := tview.NewImage()
	image.SetImage(classes_icons[(*player)["class"].(string)])
	image.SetBorder(true)

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

	gridDownRight := tview.NewGrid().
		SetRows(3, 0, 0, 0, 0, 3, 0, 3, 0).
		SetColumns(0, 0)

	update := func(refreshInterval time.Duration) {
		for *running {
			time.Sleep(refreshInterval)
			app.QueueUpdateDraw(func() {
				gridDownRight.Clear()
				gridDownRight.SetTitle("<[ " + (*player)["name"].(string) + " ]>")
				gridDownRight.SetBorder(true)
				gridDownRight.AddItem(newPrimitive("Equipement :"), 0, 0, 1, 2, 0, 0, false)
				gridDownRight.AddItem(newPrimitive("\nPoints de vie : "+strconv.Itoa((*player)["hp"].(int))+" / "+strconv.Itoa((*player)["max_hp"].(int))+"[red] ♥"), 5, 0, 1, 2, 0, 0, false)
				gridDownRight.AddItem(newBarPrimitive((*player)["hp"].(int), (*player)["max_hp"].(int), "[green]", "[red]"), 6, 0, 1, 2, 0, 0, false)
				gridDownRight.AddItem(newPrimitive("\nMana : "+strconv.Itoa((*player)["mana"].(int))+" / "+strconv.Itoa((*player)["max_mana"].(int))+"[blue] ✦"), 7, 0, 1, 2, 0, 0, false)
				gridDownRight.AddItem(newBarPrimitive((*player)["mana"].(int), (*player)["max_mana"].(int), "[blue]", "[gray]"), 8, 0, 1, 2, 0, 0, false)
				gridDownRight.AddItem(casque, 1, 0, 1, 1, 0, 0, false)
				gridDownRight.AddItem(plastron, 2, 0, 1, 1, 0, 0, false)
				gridDownRight.AddItem(jambières, 3, 0, 1, 1, 0, 0, false)
				gridDownRight.AddItem(bottes, 4, 0, 1, 1, 0, 0, false)
			})
		}
	}

	Droiteflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 0, 1, false).
		AddItem(gridDownRight, 0, 1, false)
	go update(time.Second / 24)
	return Droiteflex
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
	running := true
	// Partie droite
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)

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
			running = false
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
	running := true

	// Fonction d'arrêt différé
	differedStop := func(s int) {
		time.Sleep(time.Duration(s) * time.Second)
		running = false
		app.Stop()
		running = false
	}

	// ============================Partie droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist, &running, app)

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
		buttonActivated = 5
		combattool.MonsterAttack(player, monster, monsterList, turn)
	}

	//Fonction refresh bar de vie de l'ennemi
	actuMonsterHPBar := func() {
		bar := "Vie restante : " + strconv.Itoa((*monster)["hp"].(int)) + " / " + strconv.Itoa((*monster)["max_hp"].(int)) + " [red]♥[white]"
		bar += "\n╔───────────────────────── / ▲ " + string(byte(92)) + " ─────────────────────────╗\n[white][red]"
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
				AttackMenu.AddItem(skillList[k]["name"].(string), skillList[k]["description"].(string), rune('a'+i), func() {
					if combattool.CanPlayerUseSkill(*player, k, skillList) {
						combattool.UseSkill(player, monster, k, skillList)
						actuMonsterHPBar()
						gridCenter.RemoveItem(AttackMenu)
						gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
						monsterTurn()
					} else {
						ChatBox.Clear()
						ChatBox.SetText("Vous n'avez pas assez de mana pour lancer : " + skillList[k]["name"].(string))
					}

				})
			}
			gridCenter.AddItem(AttackMenu, 2, 0, 6, 6, 0, 0, true)
			app.SetFocus(AttackButton)
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
						combattool.UseConsumable(player, monster, k, itemlist)
						actuMonsterHPBar()
						gridCenter.RemoveItem(BackpackMenu)
						gridCenter.AddItem(MonsterIcon, 2, 0, 6, 6, 0, 0, false)
						monsterTurn()
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
	}
}

func GameOverWindow(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
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

func VictoryWindow(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
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
