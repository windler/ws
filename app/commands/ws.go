package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func GetWsDirs(root string) []string {
	return getWsDirs(root, false)
}

func GetCurrentWorkspace(root string) string {
	ws := getWsDirs(root, true)
	if len(ws) == 1 {
		return ws[0]
	}
	return ""
}

func GetWorkspaceByPattern(root, pattern string) string {
	ws := getWsDirs(root, false)
	for _, w := range ws {
		if match, _ := regexp.MatchString(pattern, w); match {
			return w
		}
	}
	return ""
}

func getWsDirs(root string, onlyCurrent bool) []string {
	result := []string{}

	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		fmt.Println(err.Error())
		return result
	}

	wd, _ := os.Getwd()
	for _, dir := range fileInfo {
		fullDir := (root + dir.Name())
		if !onlyCurrent || strings.HasPrefix(wd, fullDir) {
			result = append(result, fullDir)
		}
	}

	return result
}
