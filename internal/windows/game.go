package windows

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"

	"github.com/levor/seaBattle/internal/types"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/levor/seaBattle/internal/workers"
)

var totalTieam1 = 0
var totalTieam2 = 0

func (gw Window) CreateGameWindow() {
	label1 := widget.NewLabel("Выберите предмет")
	subjects, err := workers.GetData()
	cont := container.NewVBox()
	if err != nil {
		log.Println(err)
	}
	if len(subjects) == 0 {
		label1.Text = "Список предмеов пуст, запустите редактор и добавтье предме, тему игры и вопросы к ней!"
		btn := container.NewCenter(widget.NewButton("Редактор", gw.openEditor))
		cont = container.NewVBox(label1, btn)
	} else {
		subjectsNames := make([]string, 0)
		themeMap := make(map[string][]types.Theme, 0)
		for _, subject := range subjects {
			subjectsNames = append(subjectsNames, subject.NameSubject)
			themeMap[subject.NameSubject] = subject.Themes
		}
		selector1 := widget.NewSelect(subjectsNames, func(s string) {
			gw.chooseTheme(s, themeMap)
		})
		cont = container.NewVBox(label1, selector1)
	}

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
	maincont := container.NewVBox()
	content := container.NewHBox()
	gameField := container.NewVBox(container.New(layout.NewGridWrapLayout(fyne.NewSize(600, 2))))
	label1 := widget.NewLabel(fmt.Sprintf("Счет: %d", totalTieam1))
	label2 := widget.NewLabel(fmt.Sprintf("Счет: %d", totalTieam2))
	team1 := container.NewVBox(widget.NewLabel("Команда 1"), widget.NewSeparator(), label1)
	team2 := container.NewVBox(widget.NewLabel("Команда 2"), widget.NewSeparator(), label2)
	btnList := make([]*widget.Button, 0)
	for i := 0; i < len(questionList); i++ {
		k := i
		btn := widget.NewButton("", func() {
			gw.openDialog(questionList[k], k, label1, label2)
			btnList[k].Disable()
		})
		btnList = append(btnList, btn)
	}
	for i := 0; i < len(btnList); {
		row := container.New(layout.NewGridLayoutWithColumns(int(colColumn)))
		for j := colColumn; j > 0; j-- {
			k := i
			if k < len(btnList) {
				row.Add(btnList[k])
				i++
			} else {
				i++
			}
		}
		gameField.Add(row)
	}
	content = container.NewHBox(team1, gameField, team2)
	maincont.Add(content)
	maincont.Add(layout.NewSpacer())
	maincont.Add(widget.NewButton("Подсчитать результаты", func() {
		label1.SetText(fmt.Sprintf("Счет: %d", totalTieam1))
		label2.SetText(fmt.Sprintf("Счет: %d", totalTieam2))
	}))
	gw.cw.SetContent(maincont)
	gw.cw.Show()
}

func (gw Window) openDialog(question types.Question, k int, label1 *widget.Label, label2 *widget.Label) {
	formItems := make([]*widget.FormItem, 0)
	formItems = append(formItems, widget.NewFormItem("", widget.NewLabel(question.Question)))
	dialog := dialog.NewForm(fmt.Sprintf("Вопрос № %d", k+1), "Команда 2", "Команда 1", formItems,
		func(b bool) {
			if b {
				totalTieam2 = totalTieam2 + question.Point
			} else {
				totalTieam1 = totalTieam1 + question.Point
			}
		},
		gw.cw)
	label1.SetText(fmt.Sprintf("Счет: %d", totalTieam1))
	label2.SetText(fmt.Sprintf("Счет: %d", totalTieam2))
	dialog.Show()
}
