package main

import (
	"log/slog"
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
	logger := slog.New(slog.NewJSONHandler(streams.ErrOut, nil))
	// logger := slog.New(slog.NewTextHandler(streams.ErrOut(), nil))
	slog.SetDefault(logger)

	viperInstance := viper.New()

	cobra.OnInitialize(func() {
		// automatically read in environment variables that match supported flags
		// e.g. kubeconfig is a recognized flag so the corresponding env variable is KUBECONFIG
		viperInstance.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
		viperInstance.AutomaticEnv()
	})

	bsCmd := cmd.NewRootCmd(*common.NewHelpers(streams, viperInstance, logger))

	flags.AddFlagSet(bsCmd.PersistentFlags())
	pFlag.CommandLine = flags

	cobra.CheckErr(bsCmd.Execute())
}
