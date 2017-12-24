package contracts

import "github.com/fatih/color"

type WsInfoRetriever interface {
	Status(ws string) string
	CurrentBranch(ws string) string
}

type UI interface {
	PrintHeader(s string)
	PrintTable(header []string, rows [][]string)
	PrintString(s string, colorOrNil ...color.Attribute)
}
