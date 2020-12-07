package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"

  downloader "github.com/ralindos/coomerhelper/downloader"
)

func main() {
	a := app.New()
	w := a.NewWindow("COOOOOOOOOOOOOOOOOOM")

	hello := widget.NewLabel("KLIKNIJ TO BO NIE WYTRZYMAM")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("COOOOOM", func() {
			downloader.HandleImages()
		}),
	))

	w.ShowAndRun()
}


