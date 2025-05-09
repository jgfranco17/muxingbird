package cmds

import (
	"context"
	"io"

	"github.com/jgfranco17/muxingbird/service"
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
