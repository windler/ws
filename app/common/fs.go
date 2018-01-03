package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"
)

//EnsureDirFormat ensures that a dir has a suffix-'/' and replaces the '~' with the home dir
func EnsureDirFormat(dir string) string {
	value := strings.TrimSpace(dir)
	value = ensureHomeDir(value)
	value = ensureDirSuffic(value)

	return value
}

//EnsureFileFormat replaces the '~' with the home dir
func EnsureFileFormat(dir string) string {
	value := strings.TrimSpace(dir)
	value = ensureHomeDir(value)

	return value
}

func ensureHomeDir(dir string) string {
	result := dir
	if strings.HasPrefix(result, "~") {
		usr, _ := user.Current()
		result = strings.Replace(result, "~", usr.HomeDir, 1)
	}

	return result
}

func ensureDirSuffic(dir string) string {
	result := dir

	if !strings.HasSuffix(result, "/") {
		result = result + "/"
	}

	return result
}

func GetWsDirs(root string, onlyCurrent bool) []string {
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
