package Scene

import (
	"PPR/InventoryTool"
	"image"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreatePlayerWindow(
	classList map[string]map[string]interface{},
	icons map[string]image.Image,
	player *map[string]interface{},
	inv map[string]int,
) {
	classNames := InventoryTool.GetClassNames(classList)
	app := tview.NewApplication()
	form := tview.NewForm()
	image := tview.NewImage()

	footer := tview.NewTextView()
	footer.SetBorder(true)
	footer.SetBackgroundColor(tcell.ColorOrangeRed)
	footer.SetText("Attention, pour profité un maximum du jeux, joué dans un terminal en plein écran + mise a l'echelle de votre écran à 100% ;)")
	footer.SetTitleAlign(tview.AlignCenter)
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
			if IsNameValide(&name) {
				app.Stop()
				(*player)["name"] = name
				for k, v := range classList {
					if v["class"].(string) == class {
						(*player)["class"] = k
						(*player)["max_hp"] = v["max_hp"].(int)
						(*player)["hp"] = v["starting_hp"].(int)
						(*player)["max_mana"] = v["max_mana"].(int)
						(*player)["mana"] = v["starting_mana"].(int)
						(*player)["lvl"] = v["starting_lvl"].(int)
						tmp := []string{}
						for _, v := range v["starting_spell"].([]interface{}) {
							tmp = append(tmp, v.(string))
						}
						(*player)["skills"] = tmp

						break
					}
				}
			} else {
				app.Stop()
				CreatePlayerWindow(classList, icons, player, inv)
			}

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
	flexwithwarning := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true).
		AddItem(footer, 3, 1, false)
	if err := app.SetRoot(flexwithwarning, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

/*
Fonction qui verifie si la longeur du nom est valide et met le nom
dans le bon format (dAnIeL -> Daniel)
*/
func IsNameValide(name *string) bool {
	nom := ""
	if len(*name) > 16 || len(*name) < 3 {
		return false
	}
	for i, c := range *name {
		replace := true
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || i > 0 && c == ' ' {
			if i == 0 && c >= 'a' && c <= 'z' {
				nom += string(c - 32)
				replace = false
			} else if i > 0 && c >= 'A' && c <= 'Z' {
				nom += string(c + 32)
				replace = false
			}
		} else {
			return false
		}
		if replace {
			nom += string(c)
		}
	}
	*name = nom
	return true
}
