package commands

import "testing"
import "github.com/stretchr/testify/assert"

func TestCommandName(t *testing.T) {
	f := (&ListWsFactory{}).CreateCommand()

	assert.Equal(t, "List Workspaces", f.Name)
}
