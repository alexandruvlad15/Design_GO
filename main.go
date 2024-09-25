package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()                     //crearea aplicatiei
	w := a.NewWindow("Sum Calculator") //crearea ferestrei

	// crearea fundalului roșu
	background := canvas.NewRectangle(color.RGBA{255, 0, 0, 255})

	// crearea campurilor de introducere pt a si b
	entryA := widget.NewEntry()
	entryA.SetPlaceHolder("a=")
	entryB := widget.NewEntry()
	entryB.SetPlaceHolder("b=")
	//crearea label a si b
	labelA := canvas.NewText("Introduceti valoarea pentru a", color.RGBA{0, 128, 0, 255})
	labelB := canvas.NewText("Introduceti valoarea pentru b", color.RGBA{0, 0, 255, 255})
	// crearea unui label pentru afisarea rezultatului
	resultLabel := widget.NewLabel("")
	// crearea unui buton ce calculeaza suma
	calcButton := widget.NewButton("CALCULATE", func() {
		// convertirea valorilor din string in int
		a, _ := strconv.Atoi(entryA.Text)
		b, _ := strconv.Atoi(entryB.Text)
		sum := a + b
		resultLabel.SetText("Suma= " + strconv.Itoa(sum))
	})
	image := canvas.NewImageFromFile("/home/alex/Desktop/Designed_proj_Go/poz.jpeg")

	image.SetMinSize(fyne.NewSize(200, 200))
	// punerea tuturor widget-urilor într-un container vertical
	content := container.NewVBox(
		image,
		labelA,
		entryA,
		labelB,
		entryB,
		calcButton,
		resultLabel,
	)

	// crearea unui layout care contine fundalul si continutul
	fullContent := container.NewMax(background, content)
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(fullContent)
	w.ShowAndRun()
}
