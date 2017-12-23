package commands

import "testing"
import "github.com/stretchr/testify/assert"
import "github.com/windler/workspacehero/config"
import "github.com/windler/workspacehero/internal/test"

func TestSetupCommand(t *testing.T) {
	f := new(SetupAppFactory).CreateCommand()

	assert.Equal(t, "setup", f.Command)
	assert.Equal(t, []string{}, f.Aliases)
	assert.Equal(t, "Configure everything to unleash the beauty. Alternatively, you can edit your personal config file.", f.Description)
	assert.Equal(t, 2, len(f.Subcommands))

	scWs := f.Subcommands[0]
	assert.Equal(t, "ws", scWs.Command)
	assert.Equal(t, []string{"workspace_dir"}, scWs.Aliases)
	assert.Equal(t, "Set the root dir where all (most) of your workspaces are.", scWs.Description)

	scAdd := f.Subcommands[1]
	assert.Equal(t, "add", scAdd.Command)
	assert.Equal(t, []string{"add_single_workspace"}, scAdd.Aliases)
	assert.Equal(t, "Add an additional worskpace wich is not contained in <workspace_dir>.", scAdd.Description)
}

func TestSetNewWsDir(t *testing.T) {
	c, _ := test.CreateTestContext(config.ConfigFlag)
	repo := config.Repository(c)

	assert.NotEqual(t, "/testfile/", repo.WsDir)
	setNewWsDir(repo, "/testfile")
	assert.Equal(t, "/testfile/", repo.WsDir)
}
