package main

import (
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	pFlag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/mrlunchbox777/2024-advent-of-code/cmd"
	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

func main() {
	run(common.GetStreams)
}

func run(g common.GetStreamsFunc) {
	flags := pFlag.NewFlagSet("2024-advent-of-code", pFlag.ExitOnError)

	// setup the logger
	streams := g()

	// configure the logger
	var level slog.Leveler
	logLevel := os.Getenv("LOG_LEVEL")
	switch strings.ToLower(logLevel) {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{
		AddSource: logLevel == "debug",
		Level:     level,
	}
	logger := slog.New(slog.NewJSONHandler(streams.ErrOut, opts))
	// logger := slog.New(slog.NewTextHandler(streams.ErrOut(), nil))
	slog.SetDefault(logger)

	viperInstance := viper.New()

	cobra.OnInitialize(func() {
		// automatically read in environment variables that match supported flags
		// e.g. kubeconfig is a recognized flag so the corresponding env variable is KUBECONFIG
		viperInstance.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viperInstance.AutomaticEnv()
	})

	helpers, err := common.NewHelpers(streams, viperInstance, logger)
	cobra.CheckErr(err)
	bsCmd := cmd.NewRootCmd(helpers)

	flags.AddFlagSet(bsCmd.PersistentFlags())
	pFlag.CommandLine = flags

	cobra.CheckErr(bsCmd.Execute())
}
