package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/levor/seeBattle/internal/objects"
)

func (gw Window) CreateGameWindow() {
	w := gw.a.NewWindow("See Battle (New Game)")
	w.Resize(fyne.NewSize(800, 600))
	gw.cw = w

	// Create MainMenu
	menuItem1 := fyne.NewMenuItem("Новая игра", gw.newGame)
	menuItem2 := fyne.NewMenuItem("Редактор", gw.openEditor)
	newMenu := fyne.NewMenu("Файл", menuItem1, menuItem2)
	menu := objects.NewMainMenu(newMenu)
	w.SetMainMenu(menu.Create())

	w.SetContent(widget.NewLabel("More content"))

	w.Show()
}
