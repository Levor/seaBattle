package windows

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"

	"github.com/levor/seeBattle/internal/types"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/levor/seeBattle/internal/workers"
)

func (gw Window) CreateGameWindow() {
	label1 := widget.NewLabel("Выберите предмет")
	subjects, err := workers.GetData()
	if err != nil {
		log.Println(err)
	}
	if len(subjects) == 0 {
		fmt.Println("Empty")
	}
	subjectsNames := make([]string, 0)
	themeMap := make(map[string][]types.Theme, 0)
	for _, subject := range subjects {
		subjectsNames = append(subjectsNames, subject.NameSubject)
		themeMap[subject.NameSubject] = subject.Themes
	}
	selector1 := widget.NewSelect(subjectsNames, func(s string) {
		gw.chooseTheme(s, themeMap)
	})
	cont := container.NewVBox(label1, selector1)
	gw.cw.SetContent(cont)
	gw.cw.Show()
}

func (gw Window) chooseTheme(subName string, themeMap map[string][]types.Theme) {
	label1 := widget.NewLabel("Выбирите предмет и игра начнется автоматически")
	cont := container.NewVBox()
	if len(themeMap[subName]) == 0 {
		label1.Text = "Список тем пуст, запустите редактор и добавтье тему игры"
		btn := container.NewCenter(widget.NewButton("Редактор", gw.openEditor))
		cont = container.NewVBox(label1, btn)
	} else {
		themeList := make([]string, 0)
		questionMap := make(map[string][]types.Question)
		for _, theme := range themeMap[subName] {
			themeList = append(themeList, theme.ThemeName)
			questionMap[theme.ThemeName] = theme.Questions
		}

		selector1 := widget.NewSelect(themeList, func(s string) {
			gw.startGame(questionMap[s])
		})
		cont = container.NewVBox(label1, selector1)
	}

	gw.cw.SetContent(cont)
	gw.cw.Show()
}
func (gw Window) startGame(questionList []types.Question) {
	rand.Shuffle(len(questionList),
		func(i, j int) { questionList[i], questionList[j] = questionList[j], questionList[i] })
	colColumn := math.Ceil(math.Sqrt(float64(len(questionList))))
	content := container.NewHBox()
	gameField := container.NewVBox(container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 2))))
	team1 := container.NewVBox(widget.NewLabel("Команда 1"), widget.NewSeparator(), widget.NewLabel("Счет: 0"))
	team2 := container.NewVBox(widget.NewLabel("Команда 2"), widget.NewSeparator(), widget.NewLabel("Счет: 0"))

	for i := 0; i < len(questionList); {
		row := container.New(layout.NewGridLayoutWithColumns(int(colColumn)))
		for j := colColumn; j > 0; j-- {
			k := i
			if k < len(questionList) {
				btn := widget.NewButton("", func() {
					gw.openDialog(questionList[k], k)
				})
				row.Add(btn)
				i++
			} else {
				i++
			}
		}
		gameField.Add(row)
	}
	content = container.NewHBox(team1, gameField, team2)
	gw.cw.SetContent(content)
	gw.cw.Show()
}

func (gw Window) openDialog(question types.Question, k int) {
	dialog := dialog.NewForm(fmt.Sprintf("Вопрос № %d", k), "Команда 1", "Команда 2", nil,
		func(b bool) {
			if b {
				fmt.Println("Команда 1 заработала", question.Point)
			} else {
				fmt.Println("Команда 2 заработала", question.Point)
			}
		},
		gw.cw)
	dialog.Show()
}

func greenBTN() *fyne.Container {
	btn := widget.NewButton("", nil)
	btnColor := canvas.NewRectangle(color.NRGBA{0, 255, 0, 255})
	cont := container.New(layout.NewMaxLayout(), btn, btnColor)
	return cont
}

func redBTN() *fyne.Container {
	btn := widget.NewButton("", nil)
	btnColor := canvas.NewRectangle(color.NRGBA{255, 0, 0, 255})
	cont := container.New(layout.NewMaxLayout(), btn, btnColor)
	return cont
}

func blueBTN() *fyne.Container {
	btn := widget.NewButton("", nil)
	btnColor := canvas.NewRectangle(color.NRGBA{0, 0, 255, 255})
	cont := container.New(layout.NewMaxLayout(), btn, btnColor)
	return cont
}
