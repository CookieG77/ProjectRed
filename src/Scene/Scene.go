package Scene

import (
	"PPR/Smenu"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreatePlayerWindow() {
	app := tview.NewApplication()
	form := tview.NewForm()
	image := tview.NewImage()
	image.SetBorder(true)
	imgdata, err := Smenu.TViewMakeImg(Smenu.Icon_humain)
	imgdata2, err2 := Smenu.TViewMakeImg(Smenu.Icon_elf)
	imgdata3, err3 := Smenu.TViewMakeImg(Smenu.Icon_nain)
	if err || err2 || err3 {
		return
	}
	image.SetImage(imgdata)
	box := tview.NewTextView().SetText("[yellow]Humain:[white]\n- Points de vie: [red]100[white]\n- Points de Mana: [blue]60[white]\n\n[green]Elfe:[white]\n- Points de vie: [red]80[white]\n- Points de Mana: [blue]80[white]\n\n[purple]Nain:[white]\n- Points de vie: [red]120[white]\n- Points de Mana: [blue]40[white]\n").SetTextColor(tcell.ColorYellowGreen)
	box.SetBorder(true).SetTitle(" Stats des Classes ").SetTitleAlign(tview.AlignCenter)
	box.SetDynamicColors(true)
	form.AddDropDown("Classe", []string{"Humain", "Elfe", "Nain"}, 0, nil).
		AddInputField("Nom du Personnage", "", 16, nil, nil).
		AddButton("Jouer", func() {
			name := form.GetFormItemByLabel("Nom du Personnage").(*tview.InputField).GetText()
			_, class := form.GetFormItemByLabel("Classe").(*tview.DropDown).GetCurrentOption()
			app.Stop()
			fmt.Println(name)
			fmt.Println(class)

		}).
		AddButton("Quitter", func() {
			app.Stop()
		})

	form.GetFormItemByLabel("Classe").(*tview.DropDown).SetSelectedFunc(func(text string, index int) {
		switch index {
		case 0:
			image.SetImage(imgdata)

		case 1:
			image.SetImage(imgdata2)

		case 2:
			image.SetImage(imgdata3)
		}
	})
	form.SetBorder(true).SetTitle(" Création du personnage ").SetTitleAlign(tview.AlignCenter)
	flex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(form, 0, 1, true).
		AddItem(box, 0, 1, false).
		AddItem(image, 0, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
