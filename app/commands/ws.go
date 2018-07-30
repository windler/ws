package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type WorkspaceRetriever interface {
	GetWorkspacesIn(root string) []string
	GetCurrentWorkspace(root string) string
	GetWorkspaceByPattern(root, pattern string) string
}

type FSWorkspaceRetriever struct{}

func (wr FSWorkspaceRetriever) GetWorkspacesIn(root string) []string {
	return wr.getWsDirs(root, false)
}

func (wr FSWorkspaceRetriever) GetCurrentWorkspace(root string) string {
	ws := wr.getWsDirs(root, true)
	if len(ws) == 1 {
		return ws[0]
	}
	return ""
}

func (wr FSWorkspaceRetriever) GetWorkspaceByPattern(root, pattern string) string {
	ws := wr.getWsDirs(root, false)
	for _, w := range ws {
		if match, _ := regexp.MatchString(pattern, w); match {
			return w
		}
	}
	return ""
}

func (wr FSWorkspaceRetriever) getWsDirs(root string, onlyCurrent bool) []string {
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
