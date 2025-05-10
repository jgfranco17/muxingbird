package cmds

import (
	"github.com/jgfranco17/muxingbird/errorx"
	"github.com/jgfranco17/muxingbird/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CommandRegistry wraps the root command and core
// attributes of the CLI.
type CommandRegistry struct {
	rootCmd   *cobra.Command
	verbosity int
	logger    *logrus.Logger
}

// NewCommandRegistry creates a new instance of CommandRegistry
func NewCommandRegistry(name string, description string, version string) *CommandRegistry {
	var level logrus.Level
	logger := logging.NewLogger()
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
			logger.SetLevel(level)
			ctx := logging.ApplyToContext(cmd.Context(), logger)
			cmd.SetContext(ctx)
		},
	}
	newRegistry := &CommandRegistry{
		rootCmd: root,
		logger:  logger,
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

// Execute executes the root command and handles
// any error that may propagate. This allows us to
// return an appropriate exit code when possible.
func (cr *CommandRegistry) Execute() {
	defer errorx.HandleRecovery(cr.logger)
	err := cr.rootCmd.Execute()
	if err != nil {
		errorx.HandleError(cr.logger, err)
	}
}
