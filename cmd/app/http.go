package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"log/slog"
	"nightingale/config"
	"os"
	"os/signal"
	"syscall"
)

type Http struct {
	httpServer *echo.Echo
}

var HttpCommand = &cobra.Command{
	Use:   "http",
	Short: "Starts the HTTP server",
	Run:   RunHTTP,
}

func newHttp(httpServer *echo.Echo) Http {
	return Http{
		httpServer: httpServer,
	}
}

func RunHTTP(_ *cobra.Command, _ []string) {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	app, _, err := InitHttp(*cfg)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	msg := make(chan error)
	go func() {
		msg <- app.httpServer.Start(cfg.Server.Host + ":" + cfg.Server.Port)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		msg <- fmt.Errorf("%s", <-c)
	}()
	if err := <-msg; err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
