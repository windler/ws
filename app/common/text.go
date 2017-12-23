package common

import (
	"github.com/fatih/color"
	"github.com/windler/workspacehero/app/ui"
)

//Recommend prints a recommendation command
func Recommend(command string, ui ui.UI) {
	ui.PrintStrings([]string{"", "How about trying 'asd " + command + "'?"}, color.FgYellow)
}

//RecommendFromError prints a recommendation command after error occured
func RecommendFromError(command string, ui ui.UI) {
	ui.PrintStrings([]string{"", "Have you tried 'asd " + command + "'?"}, color.FgYellow)
}
