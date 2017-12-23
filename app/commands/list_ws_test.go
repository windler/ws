package commands

import "testing"
import "github.com/stretchr/testify/assert"

func TestLsWsCommand(t *testing.T) {
	f := new(ListWsFactory).CreateCommand()

	assert.Equal(t, "ls", f.Command)
	assert.Equal(t, []string{}, f.Aliases)
	assert.Equal(t, "List all workspaces with fancy information.", f.Description)
}
