package main

import (
	"github.com/jgfranco17/muxingbird/cmds"
	"github.com/spf13/cobra"
)

const (
	projectName        = "muxingbird"
	projectDescription = "Muxingbird: spin up configurable mock HTTP servers"
)

var (
	version string = "0.0.0-dev"
)

func main() {
	command := cmds.NewCommandRegistry(projectName, projectDescription, version)
	commandsList := []*cobra.Command{
		cmds.CommandRun(cmds.DefaultServiceFactory),
		cmds.CommandInit(),
	}
	command.RegisterCommands(commandsList)
	command.Execute()
}
