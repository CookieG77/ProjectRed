package Smenu

import (
	"PPR/InventoryTool"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	Icon_humain = "ressource/icon_humain.png"
	Icon_elf    = "ressource/icon_elf.png"
	Icon_nain   = "ressource/icon_nain.png"
	texte       = "je suis un grand tres grand meme texte qui sert a prouvé les qualité de mon ecriture et mon affichage de niveau absolu et profetionelle"
)

func SmenuRender() {
	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()

	// ============================partit droite=====================================

	image := tview.NewImage()
	imgdata, err := InventoryTool.TViewMakeImg(Icon_humain)
	if err {
		return
	}
	image.SetImage(imgdata)
	image.SetBorder(true)

	textView := tview.NewTextView().SetText(texte).SetTextColor(tcell.ColorDarkRed)
	textView.SetBorder(true)
	fmt.Print(textView, texte)

	Droiteflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(image, 0, 1, false).
		AddItem(textView, 0, 1, false)

	// =============================== partit gauche ====================================

	headergauche := tview.NewTextView().SetText("Ce deplacé a :")
	headergauche.SetBorder(true)
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)

	gaucheforet := tview.NewFlex().
		SetDirection(tview.FlexRow)

	gauchevillage := tview.NewFlex().
		SetDirection(tview.FlexRow)

	Gaucheflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gauchevillage, 0, 1, true).
		AddItem(gaucheforet, 0, 1, true)

	// =============================== partit centre ====================================

	header := tview.NewTextView().SetText("RED PROJECT ULTIMATE")
	header.SetBorder(true)
	header.SetTextColor(redColor)
	header.SetTextAlign(tview.AlignCenter)

	Centreflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false)

	// ================================ assemblage ======================================

	Machted := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(Gaucheflex, 0, 1, true).
		AddItem(Centreflex, 0, 1, true).
		AddItem(Droiteflex, 0, 1, false)

	// =============================== running =============================================

	if err4 := app.SetRoot(Machted, true).Run(); err4 != nil {
		panic(err4)
	}
}
