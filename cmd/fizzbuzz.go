package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/spf13/cobra"

	"github.com/py4mac/fizzbuzz/pkg/config"
	"github.com/py4mac/fizzbuzz/pkg/handler"
	"github.com/py4mac/fizzbuzz/pkg/server"
	"github.com/py4mac/fizzbuzz/pkg/stats/repository"
)

var (
	EndpointDefault string = ":8000"
	TimeoutDefault  int32  = 500
)

func MakeFizzBuzz() *cobra.Command {
	var serve = &cobra.Command{
		Use:          "serve",
		Short:        "Run fizzbuzz server",
		Example:      `  fizzbuzz serve`,
		SilenceUsage: true,
	}

	serve.Flags().StringP("port", "p", EndpointDefault, "Endpoint")
	serve.Flags().Int32P("timeout", "t", TimeoutDefault, "Timeout ms")

	serve.RunE = func(cmd *cobra.Command, args []string) error {
		port, _ := cmd.Flags().GetString("port")
		timeout, _ := cmd.Flags().GetInt32("timeout")

		logger.Info(
			"serve fizzbuzz",
			"Version", config.Version,
			"Built", config.Built,
			"Revision", config.Revision,
		)

		ctx, cancel := context.WithCancel(context.Background())

		s := server.NewServer(port)

		stats := repository.NewStatsInMemory()
		r := handler.NewHandler(stats, timeout)
		s.SetupHandlers(r.Engine)

		go func() {
			if err := s.Run(); err != nil && err != http.ErrServerClosed {
				logger.Fatalf("listen: %s\n", err)
			}
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		defer func() {
			signal.Stop(c)
			cancel()
		}()

		run := true

		for run {
			select {
			case <-c:
				cancel()

			case <-ctx.Done():
				run = false
			}
		}

		return nil
	}

	return serve
}
