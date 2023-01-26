package types

type Subject struct {
	NameSubject string  `json:"nameSubject"`
	Themes      []Theme `json:"themes"`
}

type Theme struct {
	ThemeName string     `json:"themeName"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Question string `json:"question"`
	Point    int    `json:"point"`
}
