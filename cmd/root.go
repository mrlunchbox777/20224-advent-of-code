package cmd

import (
	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

// NewRootCmd creates a new root command
func NewRootCmd(h common.Helpers) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "2024-advent-of-code",
		Short: "2024 Advent of Code",
		Long:  "2024 Advent of Code",
	}

	rootCmd.AddCommand(NewDay1Cmd(h))

	return rootCmd
}
