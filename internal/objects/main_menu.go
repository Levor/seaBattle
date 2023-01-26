package objects

import (
	"fyne.io/fyne/v2"
)

type MainMenu struct {
	menu *fyne.Menu
}

func NewMainMenu(menu *fyne.Menu) *MainMenu {
	return &MainMenu{menu: menu}
}

func (mm MainMenu) Create() *fyne.MainMenu {
	menu := fyne.NewMainMenu(mm.menu)
	return menu
}
