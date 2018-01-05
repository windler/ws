package ui

//PrintHeader prints a ascii art header
import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

//ConsoleUI prints output on console
type ConsoleUI struct{}

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

	seperator := []string{}
	head := []string{}
	for _, h := range header {
		seperator = append(seperator, "")
		head = append(head, strings.ToUpper("-- "+h+" --"))
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 8, 5, ' ', 0)
	fmt.Fprintln(w, strings.Join(head, "\t"))
	for _, row := range rows {
		fmt.Fprintln(w, strings.Join(row, "\t"))
	}
	w.Flush()
}
