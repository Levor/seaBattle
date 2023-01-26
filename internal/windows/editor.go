package windows

import (
	"fmt"
	"fyne.io/fyne/v2/layout"
	"github.com/levor/seeBattle/internal/workers"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/levor/seeBattle/internal/objects"
	"github.com/levor/seeBattle/internal/types"
)

func (ew Window) CreateEditorWindows() {
	ew.cw = ew.a.NewWindow("See Battle (Editor)")
	ew.cw.Resize(fyne.NewSize(800, 600))

	// Create MainMenu
	menuItem1 := fyne.NewMenuItem("Новая игра", ew.newGame)
	menuItem2 := fyne.NewMenuItem("Редактор", ew.openEditor)
	newMenu := fyne.NewMenu("Файл", menuItem1, menuItem2)
	menu := objects.NewMainMenu(newMenu)
	ew.cw.SetMainMenu(menu.Create())

	ew.createTabs()

	ew.cw.Show()
}

func (ew Window) createTabs() {
	var err error
	ew.sb, err = workers.GetData()
	if err != nil {
		log.Println(err)
	}
	tabs := container.NewAppTabs()
	for i, s := range ew.sb {
		tabcontent := container.NewVBox()
		for themid, _ := range s.Themes {
			tabcontent.Objects = append(tabcontent.Objects, ew.createThemesList(themid, i, tabcontent))
		}
		j := i

		btns := container.NewHBox(
			container.NewCenter(widget.NewButton("Добавить тему", func() {
				ew.addTheme(j)
			})),
			container.NewCenter(widget.NewButton("Удалить предмет", func() {
				ew.deleteSubject(j)
			})),
		)
		tabcontent.Add(btns)
		tabs.Items = append(tabs.Items, container.NewTabItem(s.NameSubject, tabcontent))
	}
	tabs.SetTabLocation(container.TabLocationLeading)
	container := container.NewVBox(tabs, container.NewCenter(widget.NewButton("Добавить предмет", func() {
		ew.addNewSubject()
	})))
	ew.cw.SetContent(container)
}

func (ew Window) createThemesList(themid, subid int, cont *fyne.Container) *fyne.Container {
	content := container.New(layout.NewGridLayout(3),
		widget.NewLabel(fmt.Sprintf("%d. %s", themid+1, ew.sb[subid].Themes[themid].ThemeName)),
		container.NewCenter(widget.NewButton("Редактировать", func() {
			ew.editTheme(themid, subid, cont)
		})),
		container.NewCenter(widget.NewButton("Удалить", func() {
			ew.deleteThem(themid, subid)
		})),
	)
	return content
}

func (ew Window) editTheme(themid, subid int, cont *fyne.Container) {
	cont.RemoveAll()
	label := widget.NewLabel(fmt.Sprintf("Список вопросов для темы \"%s\"", ew.sb[subid].Themes[themid].ThemeName))
	label.Alignment = fyne.TextAlignCenter
	cont.Add(label)
	entrys := make([]*widget.Entry, 0)
	for _, question := range ew.sb[subid].Themes[themid].Questions {
		entry := widget.NewEntry()
		entry.Text = question.Question
		entrys = append(entrys, entry)
	}
	btn := container.NewCenter(widget.NewButton("Добавить новый вопрос", func() {
		entrys = append(entrys, widget.NewEntry())
		ew.sb[subid].Themes[themid].Questions = append(ew.sb[subid].Themes[themid].Questions, types.Question{})
		oldContent := cont
		cont.RemoveAll()
		ew.editTheme(themid, subid, oldContent)
	}))

	btn2 := container.NewCenter(widget.NewButton("Сохранить изменения", func() {
		newQuestion := make([]types.Question, 0)
		for _, entry := range entrys {
			newQuestion = append(newQuestion, types.Question{entry.Text, 1})
		}
		ew.sb[subid].Themes[themid].Questions = newQuestion
		err := workers.UpdateData(ew.sb)
		if err != nil {
			log.Println(err)
		}
		log.Println("Data saved")
	}))
	for _, entry := range entrys {
		cont.Add(entry)
	}
	cont.Add(btn)
	cont.Add(btn2)
	cont.Add(container.NewCenter(widget.NewButton(fmt.Sprintf("Вернутся к списку тем по предмету %s", ew.sb[subid].NameSubject), func() {
		ew.createTabs()
	})))
}

func (ew Window) deleteThem(themid, subid int) {
	copy(ew.sb[subid].Themes[themid:], ew.sb[subid].Themes[subid+1:])
	ew.sb[subid].Themes[len(ew.sb[subid].Themes)-1] = types.Theme{}
	ew.sb[subid].Themes = ew.sb[subid].Themes[:len(ew.sb[subid].Themes)-1]
	err := workers.UpdateData(ew.sb)
	if err != nil {
		log.Println(err)
	}
	ew.createTabs()
}
func (ew Window) deleteSubject(subid int) {
	copy(ew.sb[subid:], ew.sb[subid+1:])
	ew.sb[len(ew.sb)-1] = types.Subject{}
	ew.sb = ew.sb[:len(ew.sb)-1]
	err := workers.UpdateData(ew.sb)
	if err != nil {
		log.Println(err)
	}
	ew.createTabs()
}

func (ew Window) addNewSubject() {
	entry := widget.NewEntry()
	ew.cw.SetContent(container.NewVBox(widget.NewLabel("Добавление нового предмета"), entry, container.NewCenter(widget.NewButton("Добавить", func() {
		if entry.Text != "" && entry.Text != "Введите текст!" {
			ew.sb = append(ew.sb, types.Subject{NameSubject: entry.Text})
			err := workers.UpdateData(ew.sb)
			if err != nil {
				log.Println(err)
			}
			ew.createTabs()
		} else {
			entry.Text = "Введите текст!"
		}
	}))))
}

func (ew Window) addTheme(sbID int) {
	entry := widget.NewEntry()
	ew.cw.SetContent(container.NewVBox(widget.NewLabel("Добавление новой темы"), entry, container.NewCenter(widget.NewButton("Добавить", func() {
		if entry.Text != "" && entry.Text != "Введите текст!" {
			ew.sb[sbID].Themes = append(ew.sb[sbID].Themes, types.Theme{ThemeName: entry.Text})
			err := workers.UpdateData(ew.sb)
			if err != nil {
				log.Println(err)
			}
			ew.createTabs()
		} else {
			entry.Text = "Введите текст!"
		}
	}))))
}
