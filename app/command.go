package app

import (
	"github.com/urfave/cli"
	"github.com/windler/ws/app/commands"
)

//CreateCliCommand creates the command for the cli-app
func CreateCliCommand(bc commands.WSCommand, cfg commands.Config) *cli.Command {
	cmd := &cli.Command{
		Name:        bc.GetCommand(),
		Usage:       bc.GetDescription(),
		Aliases:     bc.GetAliases(),
		ArgsUsage:   "[command]",
		Subcommands: createSubCommands(bc.GetSubcommands(), cfg),
		Flags:       getFlags(bc),
	}

	if bc.GetAction() != nil {
		cmd.Action = func(c *cli.Context) error {
			bc.GetAction()(commandContext{
				context: c,
				cfg:     cfg,
			})

			return nil
		}
	}

	return cmd
}

func getFlags(cmd commands.WSCommand) []cli.Flag {
	flags := []cli.Flag{}
	for _, f := range cmd.GetStringFlags() {
		switch f.GetType() {
		case "string":
			flags = append(flags, cli.StringFlag{
				Name:  f.GetName(),
				Usage: f.GetUsage(),
			})
		case "bool":
			flags = append(flags, cli.BoolFlag{
				Name:  f.GetName(),
				Usage: f.GetUsage(),
			})
		case "int":
			flags = append(flags, cli.IntFlag{
				Name:  f.GetName(),
				Usage: f.GetUsage(),
			})
		}
	}

	return flags
}

func createSubCommands(cmds []commands.WSCommand, cfg commands.Config) []cli.Command {
	subCommands := []cli.Command{}
	for _, bc := range cmds {
		command := cli.Command{
			Name:        bc.GetCommand(),
			Usage:       bc.GetDescription(),
			Aliases:     bc.GetAliases(),
			ArgsUsage:   "[command]",
			Subcommands: createSubCommands(bc.GetSubcommands(), cfg),
			Flags:       getFlags(bc),
		}

		if bc.GetAction() != nil {
			command.Action = func(c *cli.Context) error {
				bc.GetAction()(commandContext{
					context: c,
					cfg:     cfg,
				})

				return nil
			}
		}

		subCommands = append(subCommands, command)
	}
	return subCommands
}

type commandContext struct {
	context *cli.Context
	cfg     commands.Config
}

func (c commandContext) GetStringFlag(flag string) string {
	return c.context.String(flag)
}

func (c commandContext) GetBoolFlag(flag string) bool {
	return c.context.Bool(flag)
}

func (c commandContext) GetIntFlag(flag string) int {
	return c.context.Int(flag)
}

func (c commandContext) GetFirstArg() string {
	return c.context.Args().First()
}

func (c commandContext) GetConfig() commands.Config {
	return c.cfg
}
