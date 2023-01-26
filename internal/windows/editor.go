package windows

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/levor/seeBattle/internal/objects"
	"github.com/levor/seeBattle/internal/types"
)

func (ew Window) CreateEditorWindows() {
	w := ew.a.NewWindow("See Battle (Editor)")
	w.Resize(fyne.NewSize(800, 600))
	ew.cw = w

	// Create MainMenu
	menuItem1 := fyne.NewMenuItem("Новая игра", ew.newGame)
	menuItem2 := fyne.NewMenuItem("Редактор", ew.openEditor)
	newMenu := fyne.NewMenu("Файл", menuItem1, menuItem2)
	menu := objects.NewMainMenu(newMenu)
	w.SetMainMenu(menu.Create())

	fz := types.Subject{NameSubject: "Физика", Themes: []types.Theme{
		{"Физика это", []types.Question{{"Что такое физика?", 3}, {"Вопрос 1", 2}, {"Вопрос 2", 1}}},
		{"Что вы знаете о физике", []types.Question{{"Что такое физика222?", 3}, {"Вопрос 3", 2}, {"Вопрос 4", 1}}},
		{"Сила это", []types.Question{{"Что такое физика3333?", 3}, {"Вопрос 21", 2}, {"Вопрос 22", 1}}},
		{"Электричество это это", []types.Question{{"Вопрос 4", 3}}},
	}}
	math := types.Subject{NameSubject: "Математика", Themes: []types.Theme{
		{"Математика это", []types.Question{{"Что такое физика12?", 3}, {"Вопрос 1", 2}, {"Вопрос 2", 1}}},
		{"Что вы знаете о Математика", []types.Question{{"Что такое физика222?", 3}, {"Вопрос 3", 2}, {"Вопрос 4", 1}}},
		{"Считаем", []types.Question{{"Что такое физика3333?", 3}, {"Вопрос 21", 2}, {"Вопрос 22", 1}}},
		{"Рисуем", []types.Question{{"Вопрос 4", 3}}},
	}}
	subjects := make([]types.Subject, 0)
	subjects = append(subjects, fz, math)
	tab := ew.createTabs(subjects)
	w.SetContent(tab)
	w.Show()
}

func (ew Window) createTabs(subj []types.Subject) *container.AppTabs {
	tabs := container.NewAppTabs()
	for _, subject := range subj {
		tabs.Items = append(tabs.Items, ew.createTabContainer(subject))
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	return tabs
}

func (ew Window) createTabContainer(subject types.Subject) *container.TabItem {
	tabcontent := container.NewVBox()
	for i, theme := range subject.Themes {
		tabcontent.Objects = append(tabcontent.Objects, ew.createThemesList(i, theme, tabcontent, subject.NameSubject))
	}
	return container.NewTabItem(subject.NameSubject, tabcontent)
}

func (ew Window) createThemesList(i int, theme types.Theme, cont *fyne.Container, nameSubject string) *fyne.Container {
	content := container.New(layout.NewGridLayout(2),
		widget.NewLabel(fmt.Sprintf("%d. %s", i+1, theme.ThemeName)),
		widget.NewButton(fmt.Sprintf("Редактировать тему № %d", i+1), func() {
			ew.editTheme(&theme, cont, nameSubject)
		}),
	)
	return content
}

func (ew Window) editTheme(theme *types.Theme, cont *fyne.Container, nameSubject string) {
	cont.RemoveAll()
	oldContent := cont
	label := widget.NewLabel(fmt.Sprintf("Список вопросов для темы \"%s\"", theme.ThemeName))
	label.Alignment = fyne.TextAlignCenter
	cont.Add(label)
	entrys := make([]*widget.Entry, 0)
	for _, question := range theme.Questions {
		entry := widget.NewEntry()
		entry.Text = question.Question
		entrys = append(entrys, entry)
	}
	btn := widget.NewButton("Добавить новый вопрос", func() {
		entrys = append(entrys, widget.NewEntry())
		theme.Questions = append(theme.Questions, types.Question{})
		oldContent = cont
		cont.RemoveAll()
		ew.editTheme(theme, oldContent, nameSubject)
	})

	btn2 := widget.NewButton("Сохранить изминения", func() {
		newQuestion := make([]types.Question, 0)
		for _, entry := range entrys {
			newQuestion = append(newQuestion, types.Question{entry.Text, 1})
		}
		theme.Questions = newQuestion
		log.Println(theme.Questions)
		log.Println("Data saved")
	})
	for _, entry := range entrys {
		cont.Add(entry)
	}
	cont.Add(btn)
	cont.Add(btn2)
	cont.Add(widget.NewButton(fmt.Sprintf("Вернутся к списку тем по предмету %s", nameSubject), func() {
		cont.RemoveAll()
		for _, object := range oldContent.Objects {
			cont.Add(object)
		}
	}))
}
