package Smenu

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	icon_humain = "ressource/icon_humain.png"
	icon_elf    = "ressource/icon_elf.png"
	icon_nain   = "ressource/icon_nain.png"
	texte       = "je suis un grand tres grand meme texte qui sert a prouvé les qualité de mon ecriture et mon affichage de niveau absolu et profetionelle"
)

func SmenuRender() int {
	sceneValue := 0

	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()

	// ============================partit droite=====================================

	image := tview.NewImage()
	imgdata, err := TViewMakeImg(icon_humain)
	if err {
		return 5
	}
	image.SetImage(imgdata)
	image.SetBorder(true)

	textView := tview.NewTextView().SetText(texte).SetTextColor(tcell.ColorDarkRed)
	textView.SetBorder(true)
	fmt.Print(textView, texte)

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	// menu := newPrimitive("Menu")
	// main := newPrimitive("Main content")
	// sideBar := newPrimitive("Side Bar")
	casque := newPrimitive("Casque")
	plastron := newPrimitive("Plastron")
	jambières := newPrimitive("Jambières")
	bottes := newPrimitive("Bottes")

	// grid := tview.NewGrid().
	// 	SetRows(3, 0, 3).
	// 	SetColumns(0, 0, 0).
	// 	SetBorders(true).
	// 	AddItem(newPrimitive("Header"), 0, 0, 1, 3, 0, 0, false).
	// 	AddItem(newPrimitive("Footer"), 2, 0, 1, 3, 0, 0, false)

	grid := tview.NewGrid().
		SetRows(3, 0, 0, 0, 0, 3, 0, 3, 0).
		SetColumns(0, 0)
	grid.SetBorder(true)
	grid.AddItem(newPrimitive("Equipement"), 0, 0, 1, 2, 0, 0, false)
	grid.AddItem(newPrimitive("Barre de vie"), 5, 0, 1, 2, 0, 0, false)
	grid.AddItem(newPrimitive("Barre de mana"), 7, 0, 1, 2, 0, 0, false)
	// // Layout for screens narrower than 100 cells (menu and side bar are hidden).
	// grid.AddItem(menu, 0, 0, 0, 0, 0, 0, false).
	// 	AddItem(main, 1, 0, 1, 3, 0, 0, false).
	// 	AddItem(sideBar, 0, 0, 0, 0, 0, 0, false)

	grid.AddItem(casque, 1, 0, 1, 1, 0, 0, false)
	grid.AddItem(plastron, 2, 0, 1, 1, 0, 0, false)
	grid.AddItem(jambières, 3, 0, 1, 1, 0, 0, false)
	grid.AddItem(bottes, 4, 0, 1, 1, 0, 0, false)

	// // Layout for screens wider than 100 cells.
	// grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
	// 	AddItem(main, 1, 1, 1, 1, 0, 0, false).
	// 	AddItem(sideBar, 1, 2, 1, 1, 0, 0, false)

	Droiteflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 0, 1, false).
		// AddItem(textView, 0, 1, false)
		AddItem(grid, 0, 1, false)

	// =============================== partit gauche ====================================

	//header
	headergauche := tview.NewTextView().SetText("Intitié le voyage en direction de :")
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	//Buttons
	//appButtonF := tview.NewApplication()
	buttonF := tview.NewButton("go to forest").SetSelectedFunc(func() {
		sceneValue = 2
		app.Stop()
	})
	buttonF.SetBorder(true) //.SetRect(0, 0, 22, 3)

	//appButtonV := tview.NewApplication()
	buttonV := tview.NewButton("go to village").SetSelectedFunc(func() {
		sceneValue = 1
		app.Stop()
	})
	buttonV.SetBorder(true) //.SetRect(0, 0, 22, 3)

	// flex windows
	gaucheforet := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gaucheforet.SetBorder(true)
	gaucheforet.AddItem(buttonF, 3, 1, true)

	gauchevillage := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gauchevillage.SetBorder(true)
	gauchevillage.AddItem(buttonV, 3, 1, true)

	// build
	Gaucheflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headergauche, 3, 1, false).
		AddItem(gauchevillage, 0, 1, true).
		AddItem(gaucheforet, 0, 1, true)

	// =============================== partit centre ====================================

	//header
	header := tview.NewTextView().SetText("RED PROJECT ULTIMATE")
	header.SetBorder(true)
	header.SetTextColor(redColor)
	header.SetTextAlign(tview.AlignCenter)

	//build
	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false)

	// ================================ assemblage ======================================

	//build
	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Gaucheflex, 42, 1, true).
		AddItem(Centreflex, 0, 1, true).
		AddItem(Droiteflex, 0, 1, false)

	// =============================== running =============================================
	app.SetFocus(buttonV)
	if err4 := app.SetRoot(Machted, true).Run(); err4 != nil {
		return 5
		// panic(err4)
	}

	if sceneValue != 0 {
		return sceneValue
	} else {
		return 0
	}
}

// convertit les chemin d'image en crt pour un Tview image
func TViewMakeImg(addresse string) (image.Image, bool) {
	IMGbyte, err := os.ReadFile(addresse)
	graphics, err2 := png.Decode(bytes.NewReader(IMGbyte))
	if err2 != nil || err != nil {
		return nil, true
	}
	return graphics, false
}
