package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/mamad-nik/simple/apipart"
)

func helloWorld(s string) {
	a := app.New()
	w := a.NewWindow("CAT FACT")
	content := widget.NewLabel(s)
	content.Wrapping = fyne.TextWrapWord
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 10))

	//image := canvas.NewImageFromFile("/home/mamad/Downloads/catImage.jpeg")
	//image.FillMode = canvas.ImageFillOriginal
	//w.SetContent(image)
	w.ShowAndRun()
}

func main() {
	helloWorld(apipart.Runner())

}
