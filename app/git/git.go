package git

import (
	"fmt"
	"os/exec"

	"github.com/windler/asd/app/commands/contracts"
)

//Git is for all git operations
type Git struct{}

const (
	StatusCodeOk      int = 0
	StatusCodeWarning int = 1
	StatusCodeError   int = 2
)

var _ contracts.WsInfoRetriever = &Git{}

//New creates a Git Workspace Info Retriever
func New() Git {
	return Git{}
}

func (g Git) gitOnRoot(root string, args ...string) (string, error) {
	baseArgs := []string{"-C", root}
	fullArgs := append(baseArgs, args...)

	d, err := exec.Command("git", fullArgs...).Output()

	return fmt.Sprintf("%s", d), err
}
