package commands

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type WsInfoRetriever interface {
	Status(ws string) string
	CurrentBranch(ws string) string
}

func GetHeaderFunctionMap() template.FuncMap {
	return template.FuncMap{
		"wsRoot":    func(dir string) string { return "ws" },
		"gitStatus": func(dir string) string { return "git status" },
		"gitBranch": func(dir string) string { return "git branch" },
		"cmd":       func(name, dir string) string { return name },
	}
}

func GetRowsFunctionMap(infoRetriever WsInfoRetriever, markCurrentWs bool, c *WSCommandContext) template.FuncMap {
	return template.FuncMap{
		"wsRoot": func(dir string) string {
			res := dir
			wd, _ := os.Getwd()
			if markCurrentWs && strings.HasPrefix(wd, dir) {
				res = res + " <--"
			}
			return res
		},
		"gitStatus": func(dir string) string {
			return infoRetriever.Status(dir)
		},
		"gitBranch": func(dir string) string {
			return infoRetriever.CurrentBranch(dir)
		},
		"cmd": func(name, dir string) string {
			fmt.Println(dir, name)
			for _, cmd := range (*c).GetConfig().GetCustomCommands() {
				if cmd.GetName() == name {
					return strings.TrimSpace(ExecCustomCommand(&cmd, dir, c))
				}
			}
			return "-- NO OUTPUT --"
		},
	}
}
