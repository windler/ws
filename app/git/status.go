package git

//Status retrieves information wether there are untracked files
func (g Git) Status(ws string) string {
	result, err := g.gitOnRoot(ws, "status", "-s")

	if err != nil {
		return "Not a git repo"
	}

	if result == "" {
		return "UNMODIFED"
	}
	return "MODIFIED"

}
