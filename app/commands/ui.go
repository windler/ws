package commands

type UI interface {
	PrintTable(header []string, rows [][]string)
	PrintString(s string, colorOrNil ...string)
}

//Recommend prints a recommendation command
func Recommend(command string, ui UI) {
	ui.PrintString("")
	ui.PrintString("How about trying 'ws "+command+"'?", "yellow")
}

//RecommendFromError prints a recommendation command after error occured
func RecommendFromError(command string, ui UI) {
	ui.PrintString("")
	ui.PrintString("Have you tried 'ws "+command+"'?", "yellow")
}
