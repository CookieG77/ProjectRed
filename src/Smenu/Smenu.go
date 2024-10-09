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
) int {
	sceneValue := 0

	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()

	// ============================partit droite=====================================

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

	casque := newPrimitive("Casque")
	plastron := newPrimitive("Plastron")
	jambières := newPrimitive("Jambières")
	bottes := newPrimitive("Bottes")

	gridDownRight := tview.NewGrid().
		SetRows(3, 0, 0, 0, 0, 3, 0, 3, 0).
		SetColumns(0, 0)
	gridDownRight.SetTitle("<[ " + (*player)["name"].(string) + " ]>")
	gridDownRight.SetBorder(true)
	gridDownRight.AddItem(newPrimitive("Equipement :"), 0, 0, 1, 2, 0, 0, false)
	gridDownRight.AddItem(newPrimitive("\nPoints de vie: "+strconv.Itoa((*player)["hp"].(int))+" / "+strconv.Itoa((*player)["max_hp"].(int))+"[red] ♥"), 5, 0, 1, 2, 0, 0, false)
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
		sceneValue = 1
		app.Stop()
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
	gridCenter := tview.NewGrid().
		SetRows(0).
		SetColumns(0)
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
		AddItem(Gaucheflex, 42, 1, true).
		AddItem(Centreflex, 0, 1, true).
		AddItem(Droiteflex, 0, 1, false)

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
