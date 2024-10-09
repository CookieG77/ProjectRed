package Scene

import (
	"PPR/InventoryTool"
	"fmt"
	"image"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreatePlayerWindow(
	classList map[string]map[string]interface{},
	icons map[string]image.Image,
) {
	classNames := InventoryTool.GetClassNames(classList)
	app := tview.NewApplication()
	form := tview.NewForm()
	image := tview.NewImage()
	image.SetBorder(true)
	for k, v := range classList {
		if v["class"].(string) == classNames[0] {
			image.SetImage(icons[k])
			break
		}
	}
	box := tview.NewTextView().SetText(InventoryTool.MakeClassDescs(classList)).SetTextColor(tcell.ColorYellowGreen)
	box.SetBorder(true).SetTitle(" Stats des Classes ").SetTitleAlign(tview.AlignCenter)
	box.SetDynamicColors(true)
	form.AddDropDown("Classe", classNames, 0, nil).
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
		for k, v := range classList {
			if v["class"].(string) == text {
				image.SetImage(icons[k])
				return
			}
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
