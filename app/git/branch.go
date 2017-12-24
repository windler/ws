package git

import "strings"

//CurrentBranch return the current branch of the git repo
func (g Git) CurrentBranch(ws string) string {
	result, err := g.gitOnRoot(ws, "name-rev", "--name-only", "HEAD")

	if err != nil {
		return "/"
	}

	return strings.TrimSpace(result)
}
