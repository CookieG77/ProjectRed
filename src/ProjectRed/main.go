package main

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	icon_humain = "ressource/icon_humain.png"
)

func main() {
	redColor := tcell.NewRGBColor(255, 0, 0)

	IMGbyte, err := os.ReadFile(icon_humain)
	if err != nil {
		panic(err)
	}

	app := tview.NewApplication()
	image := tview.NewImage()
	//print(string(IMGbyte))
	b, err2 := base64.StdEncoding.DecodeString(string(IMGbyte[:]))
	if err2 != nil {
		panic(err2)
	}

	graphics, err3 := png.Decode(bytes.NewReader(b))
	if err3 != nil {
		panic(err3)
	}

	image.SetImage(graphics)

	box := tview.NewBox().SetBorder(true).SetTitle("[ RED_PROJECT_ULTIMATE ]")
	box.SetBorderColor(redColor)

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
