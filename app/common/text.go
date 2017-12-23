package common

import (
	"fmt"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

//PrintHeader prints a ascii art header
func PrintHeader(s string) {
	figure.NewFigure(s, "", true).Print()
	fmt.Println("")
}

//Recommend prints a recommendation command
func Recommend(command string) {
	fmt.Println("")
	color.Yellow("How about trying 'asd " + command + "'?")
}

//RecommendFromError prints a recommendation command after error occured
func RecommendFromError(command string) {
	fmt.Println("")
	color.Yellow("Have you tried 'asd " + command + "'?")
}
