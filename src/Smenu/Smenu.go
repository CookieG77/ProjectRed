package Smenu

import (
	"bytes"
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

	// Affichage de l'image
	image := tview.NewImage()
	imgdata, err := TViewMakeImg(icon_humain)
	if err {
		return 5
	}
	image.SetImage(imgdata)
	image.SetBorder(true)

	// Affichage menu equipements et barres de vie et de mana
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	casque := newPrimitive("Casque")
	plastron := newPrimitive("Plastron")
	jambières := newPrimitive("Jambières")
	bottes := newPrimitive("Bottes")

	gridDownRight := tview.NewGrid().
		SetRows(3, 0, 0, 0, 0, 3, 0, 3, 0).
		SetColumns(0, 0)
	gridDownRight.SetBorder(true)
	gridDownRight.AddItem(newPrimitive("Equipement"), 0, 0, 1, 2, 0, 0, false)
	gridDownRight.AddItem(newPrimitive("Barre de vie"), 5, 0, 1, 2, 0, 0, false)
	gridDownRight.AddItem(newPrimitive("Barre de mana"), 7, 0, 1, 2, 0, 0, false)
	gridDownRight.AddItem(casque, 1, 0, 1, 1, 0, 0, false)
	gridDownRight.AddItem(plastron, 2, 0, 1, 1, 0, 0, false)
	gridDownRight.AddItem(jambières, 3, 0, 1, 1, 0, 0, false)
	gridDownRight.AddItem(bottes, 4, 0, 1, 1, 0, 0, false)

	Droiteflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 0, 1, false).
		AddItem(gridDownRight, 0, 1, false)

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
