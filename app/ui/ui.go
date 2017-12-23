package ui

import (
	"sync"

	"github.com/fatih/color"
)

var (
	ui   UI
	once sync.Once
)

//UI handles output
type UI interface {
	PrintHeader(s string)
	PrintTable(header []string, rows [][]string)
	PrintString(s string, colorOrNil ...color.Attribute)
	PrintStrings(s []string, colorOrNil ...color.Attribute)
}

//CurrentUI returns the current used ui
func CurrentUI() UI {
	return ui
}

//SetUI sets the ui for the app
func SetUI(u UI) {
	once.Do(func() {
		if ui == nil {
			ui = u
		}
	})
}
