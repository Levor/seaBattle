package windows

import (
	"github.com/levor/seeBattle/internal/types"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/levor/seeBattle/internal/objects"
)

type Window struct {
	a  fyne.App
	cw fyne.Window
	sb []types.Subject
}

func NewHomeWindow(a fyne.App) *Window {
	return &Window{a: a}
}

func (hw Window) CreateHomeWindow() {
	w := hw.a.NewWindow("See Battle")
	w.Resize(fyne.NewSize(800, 600))
	hw.cw = w

	// Create MainMenu
	menuItem1 := fyne.NewMenuItem("Новая игра", hw.newGame)
	menuItem2 := fyne.NewMenuItem("Редактор", hw.openEditor)
	newMenu := fyne.NewMenu("Файл", menuItem1, menuItem2)
	menu := objects.NewMainMenu(newMenu)
	w.SetMainMenu(menu.Create())

	// label = "Добро пожаловать в интерактивную интеллектуальную игру "Морской бой" "
	text := canvas.NewText("Добро пожаловать", color.RGBA{0, 50, 200, 10})
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle = fyne.TextStyle{Bold: true}
	text.TextSize = 20
	text2 := canvas.NewText("в интерактивную интеллектуальную игру", color.RGBA{0, 50, 200, 10})
	text2.Alignment = fyne.TextAlignCenter
	text2.TextStyle = fyne.TextStyle{Bold: true}
	text2.TextSize = 20
	text3 := canvas.NewText("\"Морской бой\"", color.RGBA{0, 50, 200, 10})
	text3.Alignment = fyne.TextAlignCenter
	text3.TextStyle = fyne.TextStyle{Bold: true}
	text3.TextSize = 20

	btnNewGame := widget.NewButton("Начать игру", hw.newGame)
	separator1 := widget.NewSeparator()
	separator2 := widget.NewSeparator()

	img := canvas.NewImageFromFile("seeBattle.png")
	img.SetMinSize(fyne.NewSize(400, 300))
	imgContainer := container.NewCenter(img)
	content := container.NewVBox(text, text2, text3, separator1, imgContainer, separator2, container.NewCenter(btnNewGame))
	w.SetContent(content)
	w.Show()
}

func (hw Window) newGame() {
	hw.cw.Close()
	hw.CreateGameWindow()
	log.Println("New game open successfully")
}

func (hw Window) openEditor() {
	hw.cw.Close()
	hw.CreateEditorWindows()
	log.Println("Editor open successfully")
}
