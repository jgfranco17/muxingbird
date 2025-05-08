package cmds

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/jgfranco17/muxingbird/logging"
	"github.com/jgfranco17/muxingbird/service"
	"github.com/spf13/cobra"
)

const (
	defaultPort        int           = 8000
	defaultMaxDuration time.Duration = 5 * time.Minute
)

type HttpService interface {
	Run(ctx context.Context) error
}

type ShellExecutor interface {
	Exec(ctx context.Context, name string, args string) (int, string, error)
}

type ServiceFactory func(ctx context.Context, r io.Reader, port int) (HttpService, error)

func DefaultServiceFactory(ctx context.Context, r io.Reader, port int) (HttpService, error) {
	srv, err := service.NewRestService(ctx, r, port)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

// CommandRun creates a new Cobra command for running the HTTP service.
func CommandRun(serviceFactory ServiceFactory) *cobra.Command {
	var port int
	var activeDuration time.Duration
	cmd := &cobra.Command{
		Use:   "run",
		Args:  cobra.ExactArgs(1),
		Short: "Run the server from the config",
		Long:  "Spin up the HTTP service based on the definitions file",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logging.FromContext(cmd.Context())
			ctx, cancel := context.WithTimeout(cmd.Context(), activeDuration)
			defer cancel()
			path := args[0]
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			logger.Debugf("Using config: %s", path)
			server, err := serviceFactory(ctx, file, port)
			if err != nil {
				return err
			}
			logger.Debugf("Server configured with uptime of %s", activeDuration)
			return server.Run(ctx)
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.Flags().DurationVarP(&activeDuration, "duration", "d", defaultMaxDuration, "Maximum duration to run server")
	cmd.Flags().IntVarP(&port, "port", "p", defaultPort, "Port to run server on")
	return cmd
}
