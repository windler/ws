package git

//Status represents the git repo status
type Status struct {
	Status string
	Code   int
}

//Status retrieves information wether there are untracked files
func (g Git) Status() Status {
	result, err := g.gitOnRoot("status", "-s")

	if err != nil {
		return Status{
			Status: "Not a git repo",
			Code:   StatusCodeWarning,
		}
	}

	if result == "" {
		return Status{
			Status: "UNMODIFED",
			Code:   StatusCodeOk,
		}
	}
	return Status{
		Status: "MODIFIED",
		Code:   StatusCodeError,
	}
}
