package Smenu

import (
	"image"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func SmenuRender(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
) int {
	sceneValue := 0

	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()

	// ============================partit droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)

	// =============================== partit gauche ====================================

	// ++++++ header ++++++
	headergauche := tview.NewTextView().SetText("Intitié le voyage en direction de :")
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	// ++++++ images ++++++
	// imageV := tview.NewImage()
	// imgdataV, err := TViewMakeImg(icon_village)
	// if err {
	// 	return 5
	// }
	// imageV.SetImage(imgdataV)
	// imageV.SetBorder(true)

	// imageF := tview.NewImage()
	// imgdataF, err := TViewMakeImg(icon_forest)
	// if err {
	// 	return 5
	// }
	// imageF.SetImage(imgdataF)
	// imageF.SetBorder(true)

	// ++++++ Buttons ++++++
	buttonF := tview.NewButton("go to forest").SetSelectedFunc(func() {
		sceneValue = 2
		app.Stop()
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
	//gaucheforet.AddItem(imageF, 0, 1, false)

	gauchevillage := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gauchevillage.SetBorder(true)
	gauchevillage.AddItem(buttonV, 3, 1, true)
	//gauchevillage.AddItem(imageV, 0, 1, false)

	// ++++++ build ++++++
	Gaucheflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headergauche, 3, 1, true).
		AddItem(gauchevillage, 0, 1, true).
		AddItem(gaucheforet, 0, 1, true)

	// =============================== partit centre ====================================

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
		ShowInventory(classes_icons, bg_imgs, player, itemlist, inv)
	})
	invBoutton.SetBorder(true)
	gridCenter := tview.NewGrid().
		SetRows(0, 0, 0, 0).
		SetColumns(0, 0, 0, 0).
		AddItem(quitButton, 4, 4, 1, 1, 0, 0, true).
		AddItem(invBoutton, 4, 2, 1, 2, 0, 0, true)
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
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Buttons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			SmenuRender(classes_icons, bg_imgs, player, itemlist, inv)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)

	consumableButton := tview.NewButton("Consommable").
		SetSelectedFunc(func() {
			app.Stop()
			ShowConsumable(classes_icons, bg_imgs, player, itemlist, inv)
		})

	equipementButton := tview.NewButton("Equipement").
		SetSelectedFunc(func() {
			app.Stop()
			ShowEquipement(classes_icons, bg_imgs, player, itemlist, inv)
		})
	othersButton := tview.NewButton("Autre").
		SetSelectedFunc(func() {
			app.Stop()
			ShowOthers(classes_icons, bg_imgs, player, itemlist, inv)
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
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Buttons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, player, itemlist, inv)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	app.Stop()

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
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
) {
	app := tview.NewApplication()
	// ============================partit droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Buttons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, player, itemlist, inv)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	app.Stop()

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
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
) int {
	app := tview.NewApplication()
	// ============================partit droite=====================================
	Droiteflex := CreateRightPart(classes_icons, player, itemlist)
	// Buttons
	quitButton := tview.NewButton("Retour").
		SetSelectedFunc(func() {
			app.Stop()
			ShowInventory(classes_icons, bg_imgs, player, itemlist, inv)
		})
	quitButton.SetBackgroundColor(tcell.ColorRed)
	// TextView for inv items
	app.Stop()

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

	Droiteflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 0, 1, false).
		AddItem(gridDownRight, 0, 1, false)
	return Droiteflex
}

func CreateLeftPart(
	classes_icons map[string]image.Image,
	bg_imgs map[string]image.Image,
	player *map[string]interface{},
	itemlist map[string]map[string]interface{},
	inv map[string]int,
	imageTop map[string]image.Image,
	imageBottom map[string]image.Image,
	directionTop tview.Primitive,
	directionBottom tview.Primitive,
	nameTop string,
	nameBottom string,
	headerText string,
	app *tview.Application,
) *tview.Flex {

	headergauche := tview.NewTextView().SetText(headerText)
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	buttonTOP := tview.NewButton(nameTop).SetSelectedFunc(func() {
		app.Stop()
		if err4 := app.SetRoot(directionTop, true).EnableMouse(true).Run(); err4 != nil {
			panic(err4)
		}
		//forgeron
	})
	buttonTOP.SetBorder(true)

	buttonBottom := tview.NewButton(nameBottom).SetSelectedFunc(func() {
		app.Stop()
		if err4 := app.SetRoot(directionBottom, true).EnableMouse(true).Run(); err4 != nil {
			panic(err4)
		}
		//vendeur
	})
	buttonBottom.SetBorder(true)

	flexTOP := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(buttonBottom, 3, 1, true)
		// futur une image
	flexTOP.SetBorder(true)

	flexBottom := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(buttonBottom, 3, 1, true)
		// futur une image
	flexBottom.SetBorder(true)

	GaucheFlex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headergauche, 3, 1, true).
		AddItem(flexTOP, 0, 1, true).
		AddItem(flexBottom, 0, 1, true)

	return GaucheFlex
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

	linkTop := Sshopvillage(classes_icons, bg_imgs, player, itemlist, inv, "forgeron")
	linkBottom := Sshopvillage(classes_icons, bg_imgs, player, itemlist, inv, "vendeur")

	Gaucheflex := CreateLeftPart(classes_icons, bg_imgs, player, itemlist, inv, bg_imgs, bg_imgs, linkTop, linkBottom,
		"top", "bottom", "Rendre visite a votre tres cher :", app)

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
) tview.Primitive {
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
		textTopgauche.SetText("Hmfmfhmfm, bricole cassé, objet a bricolé, bienvenu dans les flamme et le fer, rien de mieux pour battre l'enfer mfmfmfm")

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

	return Machted
	// if err4 := app.SetRoot(Machted, true).EnableMouse(true).Run(); err4 != nil {
	// 	panic(err4)
	// }
}
