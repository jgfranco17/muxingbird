package main

import (
	"github.com/jgfranco17/muxingbird/cmds"
	"github.com/jgfranco17/muxingbird/logging"
	"github.com/spf13/cobra"
)

const (
	projectName        = "muxingbird"
	projectDescription = "Muxingbird: spin up HTTP servers with a few clicks"
)

var (
	version string = "0.0.0-dev"
)

func main() {
	logger := logging.NewLogger()
	command := cmds.NewCommandRegistry(projectName, projectDescription, version)
	commandsList := []*cobra.Command{
		cmds.CommandRun(cmds.DefaultServiceFactory),
	}
	command.RegisterCommands(commandsList)
	if err := command.Execute(); err != nil {
		logger.Fatalf("Muxingbird error: %v", err)
	}
}
