package ui

//PrintHeader prints a ascii art header
import (
	"fmt"
	"os"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

//ConsoleUI prints output on console
type ConsoleUI struct{}

func (ui ConsoleUI) PrintHeader(s string) {
	figure.NewFigure(s, "", true).Print()
	fmt.Println("")
}

func (ui ConsoleUI) PrintString(s string, colorOrNil ...string) {
	var attr color.Attribute

	if len(colorOrNil) == 1 {

		switch colorOrNil[0] {
		case "white":
			attr = color.FgWhite

		case "green":
			attr = color.FgGreen

		case "yellow":
			attr = color.FgYellow

		case "red":
			attr = color.FgRed
		}
	}

	color.New(attr).Println(s)
}

func (ui ConsoleUI) PrintTable(header []string, rows [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorder(false)
	table.AppendBulk(rows)
	table.Render()
}
