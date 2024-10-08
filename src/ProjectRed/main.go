package main

import (
<<<<<<< Updated upstream
=======
	"PPR"
	"bytes"
	"image"
	"image/png"
>>>>>>> Stashed changes
	"os"
)

func main() {
<<<<<<< Updated upstream
=======
	redColor := tcell.NewRGBColor(255, 0, 0)
	PPR.FormTest()
	app := tview.NewApplication()
	image := tview.NewImage()
	box := tview.NewBox().SetBorder(true).SetTitle("[ RED PROJECT ULTIMATE ]")
	box.SetBorderColor(redColor)
	imgdata, err := TViewMakeImg(icon_nain)
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

>>>>>>> Stashed changes
}

func getArgs() []string {
	return os.Args[1:]
}
