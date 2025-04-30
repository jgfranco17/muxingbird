package cmds

import (
	"context"
	"os"

	"github.com/jgfranco17/muxingbird/logging"
	"github.com/jgfranco17/muxingbird/service"
	"github.com/spf13/cobra"
)

type ShellExecutor interface {
	Exec(ctx context.Context, name string, args string) (int, string, error)
}

func CommandRun() *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use:   "run",
		Args:  cobra.ExactArgs(1),
		Short: "Run the server from the config",
		Long:  "Spin up the HTTP service based on the definitions file",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logging.NewLogger()
			path := args[0]
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			server, err := service.NewRestService(file, port)
			if err != nil {
				return err
			}
			logger.Infof("Starting service on port %d", port)
			return server.Run()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.Flags().IntVarP(&port, "port", "p", 8000, "Port to run server on")
	return cmd
}
