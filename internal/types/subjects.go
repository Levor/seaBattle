package types

type Subject struct {
	NameSubject string
	Themes      []Theme
}

type Theme struct {
	ThemeName string
	Questions []Question
}

type Question struct {
	Question string
	Point    int
}
