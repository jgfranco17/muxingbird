package cmds

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbosity int
)

type CommandRegistry struct {
	rootCmd   *cobra.Command
	verbosity int
	logger    *logrus.Logger
}

// NewCommandRegistry creates a new instance of CommandRegistry
func NewCommandRegistry(name string, description string, version string) *CommandRegistry {
	var level logrus.Level
	root := &cobra.Command{
		Use:     name,
		Version: version,
		Short:   description,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			verbosityFlagCount, _ := cmd.Flags().GetCount("verbose")
			switch verbosityFlagCount {
			case 1:
				level = logrus.InfoLevel
			case 2:
				level = logrus.DebugLevel
			case 3:
				level = logrus.TraceLevel
			default:
				level = logrus.WarnLevel
			}
			logrus.SetLevel(level)
		},
	}
	newRegistry := &CommandRegistry{
		rootCmd: root,
	}
	root.PersistentFlags().CountVarP(&newRegistry.verbosity, "verbose", "v", "Increase verbosity (-v or -vv)")
	return newRegistry
}

// RegisterCommand registers a new command within the CommandRegistry
func (cr *CommandRegistry) RegisterCommands(commands []*cobra.Command) {
	for _, cmd := range commands {
		cr.rootCmd.AddCommand(cmd)
	}
}

// Execute executes the root command
func (cr *CommandRegistry) Execute() error {
	return cr.rootCmd.Execute()
}
