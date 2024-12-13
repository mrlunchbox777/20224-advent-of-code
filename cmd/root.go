package cmd

import (
	"github.com/mrlunchbox777/2024-advent-of-code/cmd/day1"
	"github.com/mrlunchbox777/2024-advent-of-code/cmd/day2"
	"github.com/mrlunchbox777/2024-advent-of-code/cmd/day3"
	"github.com/mrlunchbox777/2024-advent-of-code/cmd/day4"
	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

// NewRootCmd creates a new root command
func NewRootCmd(h *common.Helpers) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "2024-advent-of-code",
		Short: "2024 Advent of Code",
		Long:  "2024 Advent of Code",
	}

	rootCmd.AddCommand(day1.NewCmd(h))
	rootCmd.AddCommand(day2.NewCmd(h))
	rootCmd.AddCommand(day3.NewCmd(h))
	rootCmd.AddCommand(day4.NewCmd(h))

	return rootCmd
}
