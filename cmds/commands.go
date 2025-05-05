package cmds

import (
	"context"
	"io"
	"os"

	"github.com/jgfranco17/muxingbird/service"
	"github.com/spf13/cobra"
)

type HttpService interface {
	Run() error
}

type ShellExecutor interface {
	Exec(ctx context.Context, name string, args string) (int, string, error)
}

type ServiceFactory func(r io.Reader, port int) (HttpService, error)

func DefaultServiceFactory(r io.Reader, port int) (HttpService, error) {
	srv, err := service.NewRestService(r, port)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

func CommandRun(serviceFactory ServiceFactory) *cobra.Command {
	var port int
	cmd := &cobra.Command{
		Use:   "run",
		Args:  cobra.ExactArgs(1),
		Short: "Run the server from the config",
		Long:  "Spin up the HTTP service based on the definitions file",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			server, err := serviceFactory(file, port)
			if err != nil {
				return err
			}
			return server.Run()
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.Flags().IntVarP(&port, "port", "p", 8000, "Port to run server on")
	return cmd
}
