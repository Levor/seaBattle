package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/levor/seaBattle/internal/windows"
	"io/ioutil"
)

func main() {
	a := app.New()
	file, _ := ioutil.ReadFile("logo.png")
	a.SetIcon(fyne.NewStaticResource("icon", file))
	game := windows.NewHomeWindow(a)
	game.CreateHomeWindow()
	a.Run()
}
