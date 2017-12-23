package commands

import (
	"bufio"
	"fmt"

	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"github.com/windler/workspacehero/app/common"
	"github.com/windler/workspacehero/config"
)

//SetupAppFactory creates commands to list workspace information
type SetupAppFactory struct{}

//ensure interface
var _ BaseCommandFactory = &SetupAppFactory{}

//CreateCommand creates a ListWsCommand
func (factory *SetupAppFactory) CreateCommand() BaseCommand {

	setWsDirSubCommand := BaseCommand{
		Command:     "ws",
		Description: "Set the root dir where all (most) of your workspaces are.",
		Aliases:     []string{"workspace_dir"},
		Action:      setWsDirSubCommandExec,
		Subcommands: []BaseCommand{},
	}

	addWsSubCommand := BaseCommand{
		Command:     "add",
		Description: "Add an additional worskpace wich is not contained in <workspace_dir>.",
		Aliases:     []string{"add_single_workspace"},
		Action:      setAddWsSubCommandExec,
		Subcommands: []BaseCommand{},
	}

	return BaseCommand{
		Command:     CmdSetup,
		Description: "Configure everything to unleash the beauty. Alternatively, you can edit your personal config file.",
		Aliases:     []string{},
		Subcommands: []BaseCommand{
			setWsDirSubCommand,
			addWsSubCommand,
		},
	}
}

func setWsDirSubCommandExec(c *cli.Context) error {
	repo := config.Repository(c)

	common.PrintHeader("Workspace")

	color.White("Current workspace dir to scan: ")
	color.Green(repo.WsDir)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	color.White("New value: ")
	newWsDir, _ := reader.ReadString('\n')
	fmt.Println("")

	setNewWsDir(repo, newWsDir)

	return nil
}

func setNewWsDir(repo *config.Config, dir string) {
	repo.WsDir = common.EnsureDirFormat(dir)
	repo.Save()

	color.White("Successfully set to:")
	color.Green(repo.WsDir)

	common.Recommend(CmdListWs)
}

func setAddWsSubCommandExec(c *cli.Context) error {

	return nil
}
