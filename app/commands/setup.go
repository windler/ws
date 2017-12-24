package commands

import (
	"bufio"

	"github.com/windler/workspacehero/app/commands/contracts"

	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"github.com/windler/workspacehero/app/common"
	"github.com/windler/workspacehero/config"
)

//SetupAppFactory creates commands to list workspace information
type SetupAppFactory struct {
	UserInterface contracts.UI
}

//ensure interface
var _ BaseCommandFactory = &SetupAppFactory{}

//CreateCommand creates a ListWsCommand
func (factory *SetupAppFactory) CreateCommand() BaseCommand {

	setWsDirSubCommand := BaseCommand{
		Command:     "ws",
		Description: "Set the root dir where all (most) of your workspaces are.",
		Aliases:     []string{"workspace_dir"},
		Action: func(c *cli.Context) error {
			return factory.setWsDirSubCommandExec(c)
		},
		Subcommands: []BaseCommand{},
	}

	addWsSubCommand := BaseCommand{
		Command:     "add",
		Description: "Add an additional worskpace wich is not contained in <workspace_dir>.",
		Aliases:     []string{"add_single_workspace"},
		Action: func(c *cli.Context) error {
			return factory.addWsSubCommandExec(c)
		},
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

func (factory *SetupAppFactory) UI() contracts.UI {
	return factory.UserInterface
}

func (factory *SetupAppFactory) setWsDirSubCommandExec(c *cli.Context) error {
	repo := config.Repository(c)

	factory.UI().PrintHeader("Workspace")

	factory.UI().PrintString("Current workspace dir to scan: ")
	factory.UI().PrintString(repo.WsDir, color.FgGreen)

	reader := bufio.NewReader(os.Stdin)
	factory.UI().PrintString("New value: ")
	newWsDir, _ := reader.ReadString('\n')
	factory.setNewWsDir(repo, newWsDir)

	return nil
}

func (factory *SetupAppFactory) setNewWsDir(repo *config.Config, dir string) {
	repo.WsDir = common.EnsureDirFormat(dir)
	repo.Save()

	factory.UI().PrintString("Successfully set to:")
	factory.UI().PrintString(repo.WsDir, color.FgGreen)

	Recommend(CmdListWs, factory.UI())
}

func (factory *SetupAppFactory) addWsSubCommandExec(c *cli.Context) error {

	return nil
}
