package git

import (
	"fmt"
	"os/exec"
)

//Git is for all git operations
type Git struct {
	Root string
}

const (
	StatusCodeOk      int = 0
	StatusCodeWarning int = 1
	StatusCodeError   int = 2
)

//For creates a Git Object for the given Directory
func For(path string) Git {
	return Git{
		Root: path,
	}
}

func (g Git) gitOnRoot(args ...string) (string, error) {
	baseArgs := []string{"-C", g.Root}
	fullArgs := append(baseArgs, args...)

	d, err := exec.Command("git", fullArgs...).Output()

	return fmt.Sprintf("%s", d), err
}
