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

func SmenuRender() {
	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()

	// ============================partit droite=====================================

	image := tview.NewImage()
	imgdata, err := TViewMakeImg(icon_humain)
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

	headergauche := tview.NewTextView().SetText("Intitié le voyage en direction de :")
	headergauche.SetBorder(true)
	headergauche.SetTextColor(tcell.ColorGhostWhite)
	headergauche.SetTextAlign(tview.AlignCenter)
	headergauche.SetBorder(true)

	appButtonF := tview.NewApplication()
	buttonF := tview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
		appButtonF.Stop()
	})
	buttonF.SetBorder(true).SetRect(0, 0, 22, 3)

	appButtonV := tview.NewApplication()
	buttonV := tview.NewButton("Hit Enter to close").SetSelectedFunc(func() {
		appButtonV.Stop()
	})
	buttonV.SetBorder(true).SetRect(0, 0, 22, 3)

	gaucheforet := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gaucheforet.SetBorder(true)
	gaucheforet.AddItem(buttonF, 3, 1, false)

	gauchevillage := tview.NewFlex().
		SetDirection(tview.FlexRow)
	gauchevillage.SetBorder(true)
	gauchevillage.AddItem(buttonV, 3, 1, false)

	Gaucheflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(headergauche, 3, 1, false).
		AddItem(gauchevillage, 0, 1, false).
		AddItem(gaucheforet, 0, 1, false)

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
		AddItem(Gaucheflex, 42, 1, false).
		AddItem(Centreflex, 0, 1, false).
		AddItem(Droiteflex, 0, 1, false)

	// =============================== running =============================================

	if err4 := app.SetRoot(Machted, true).Run(); err4 != nil {
		panic(err4)
	}
}

func TViewMakeImg(addresse string) (image.Image, bool) {
	IMGbyte, err := os.ReadFile(addresse)
	graphics, err2 := png.Decode(bytes.NewReader(IMGbyte))
	if err2 != nil || err != nil {
		return nil, true
	}
	return graphics, false
}
