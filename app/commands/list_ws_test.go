package commands

import "testing"
import "github.com/stretchr/testify/assert"

func TestCommandName(t *testing.T) {
	f := new(ListWsFactory).CreateCommand()

	assert.Equal(t, "workspace", f.Command)
	assert.Equal(t, []string{"ws"}, f.Aliases)
	assert.Equal(t, "List all workspaces with fancy information", f.Description)
}
