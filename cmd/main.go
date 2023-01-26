package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/levor/seeBattle/internal/windows"
)

func main() {
	a := app.New()
	game := windows.NewHomeWindow(a)
	game.CreateHomeWindow()
	a.Run()
}
