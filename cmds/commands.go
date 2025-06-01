package cmds

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jgfranco17/muxingbird/errorx"
	"github.com/jgfranco17/muxingbird/logging"
	"github.com/jgfranco17/muxingbird/service"
	"github.com/spf13/cobra"
)

const (
	defaultPort        int           = 8000
	defaultMaxDuration time.Duration = 5 * time.Minute
)

// CommandRun creates a new Cobra command for running the HTTP service.
func CommandRun(serviceFactory ServiceFactory) *cobra.Command {
	var configFile string
	var port int
	var activeDuration time.Duration

	cmd := &cobra.Command{
		Use:   "run",
		Args:  cobra.ExactArgs(0),
		Short: "Run the server from the config",
		Long:  "Spin up the HTTP service based on the definitions file",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logging.FromContext(cmd.Context())
			ctx, cancel := context.WithTimeout(cmd.Context(), activeDuration)
			defer cancel()
			path, err := filepath.Abs(configFile)
			if err != nil {
				return errorx.NewErrorWithCode(err, errorx.ExitInvalidArgs)
			}
			file, err := os.Open(path)
			if err != nil {
				return errorx.NewErrorWithCode(err, errorx.ExitFileNotFound)
			}
			logger.Debugf("Using config: %s", path)
			server, err := serviceFactory(ctx, file, port)
			if err != nil {
				return errorx.NewErrorWithCode(err, errorx.ExitConfigError)
			}
			logger.Debugf("Server configured on port %d with uptime of %s", port, activeDuration)
			return server.Run(ctx)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Flags().StringVarP(&configFile, "config-file", "f", service.ConfigFileName, "Service definition file")
	cmd.Flags().DurationVarP(&activeDuration, "duration", "d", defaultMaxDuration, "Maximum duration to run server")
	cmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Port to run server on")

	return cmd
}

// CommandInit creates a new Cobra command for initializing
// a starter config with basic routes.
func CommandInit() *cobra.Command {
	var outputPath string
	var force bool
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Generate a starter mock config file",
		Long:  "Create a basic starter server definition file for the user to fill.",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logging.FromContext(cmd.Context())
			ctx, cancel := context.WithCancel(cmd.Context())
			defer cancel()
			if outputPath == "" {
				if err := service.InitConfig(ctx, cmd.OutOrStdout()); err != nil {
					return errorx.NewErrorWithCode(err, errorx.ExitGenericError)
				}
				return nil
			}
			if _, err := os.Stat(outputPath); err == nil && !force {
				coreErr := fmt.Errorf("file %q already exists (use --force to overwrite)", outputPath)
				return errorx.NewErrorWithCode(coreErr, errorx.ExitInvalidArgs)

			}
			file, err := os.Create(outputPath)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			defer func() {
				if cerr := file.Close(); cerr != nil {
					logger.Fatalf("Failed to close file: %v", cerr)
				}
			}()
			if err := service.InitConfig(ctx, file); err != nil {
				return errorx.NewErrorWithCode(err, errorx.ExitGenericError)
			}
			logger.Infof("Created new default config: %s", outputPath)
			return nil
		},
	}
	cmd.Flags().StringVarP(&outputPath, "output", "o", "", "Path to write config file (default: stdout)")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing file if it exists")
	return cmd
}
