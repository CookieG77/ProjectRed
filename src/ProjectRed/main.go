package main

import (
	"bytes"
	"image"
	"image/png"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	icon_humain = "ressource/icon_elf.png"
	icon_elf    = "ressource/icon_elf.png"
	icon_nain   = "ressource/icon_nain.png"
)

func main() {
	redColor := tcell.NewRGBColor(255, 0, 0)

	app := tview.NewApplication()
	image := tview.NewImage()
	box := tview.NewBox().SetBorder(true).SetTitle("[ RED_PROJECT_ULTIMATE ]")
	box.SetBorderColor(redColor)
	imgdata, err := TviewMakeImg(icon_humain)
	if err {
		return
	}
	image.SetImage(imgdata)
	if err4 := app.SetRoot(image, true).Run(); err4 != nil {
		panic(err4)
	}
	// run := true
	// for !run {
	// 	clearCmd()

	// 	time.Sleep(time.Second / 24)
	// }

}

func getArgs() []string {
	return os.Args[1:]
}

// func clearCmd() {
// 	if runtime.GOOS == "windows" {
// 		cmd := exec.Command("cmd", "/c", "cls")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	} else {
// 		cmd := exec.Command("clear")
// 		cmd.Stdout = os.Stdout
// 		cmd.Run()
// 	}
// }

func TViewMakeImg(addresse string) (image.Image, bool) {
	IMGbyte, err := os.ReadFile(icon_humain)
	graphics, err2 := png.Decode(bytes.NewReader(IMGbyte))
	if err2 != nil || err != nil {
		return nil, true
	}
	return graphics, false
}
