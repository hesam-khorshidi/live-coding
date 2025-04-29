package main

import (
	"github.com/spf13/cobra"
	"live-coding/cmd/app"
	"log/slog"
)

func main() {
	cmd := &cobra.Command{
		Use:   "live-coding",
		Short: "live code test",
		Long: `
▗▖   ▗▄▄▄▖▗▖  ▗▖▗▄▄▄▖     ▗▄▄▖ ▗▄▖ ▗▄▄▄ ▗▄▄▄▖▗▖  ▗▖ ▗▄▄▖
▐▌     █  ▐▌  ▐▌▐▌       ▐▌   ▐▌ ▐▌▐▌  █  █  ▐▛▚▖▐▌▐▌   
▐▌     █  ▐▌  ▐▌▐▛▀▀▘    ▐▌   ▐▌ ▐▌▐▌  █  █  ▐▌ ▝▜▌▐▌▝▜▌
▐▙▄▄▖▗▄█▄▖ ▝▚▞▘ ▐▙▄▄▖    ▝▚▄▄▖▝▚▄▞▘▐▙▄▄▀▗▄█▄▖▐▌  ▐▌▝▚▄▞▘
`,
	}
	cmd.AddCommand(app.HttpCommand)

	if err := cmd.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
