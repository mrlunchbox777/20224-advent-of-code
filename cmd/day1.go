package cmd

import (
	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

// NewDay1Cmd creates a new day1 command
func NewDay1Cmd(h common.Helpers) *cobra.Command {
	day1Cmd := &cobra.Command{
		Use:   "day1",
		Short: "Day 1",
		Long:  "Day 1",
		RunE: func(cmd *cobra.Command, args []string) error {
			return day1(h, cmd, args)
		},
	}

	return day1Cmd
}

func day1(h common.Helpers, cmd *cobra.Command, args []string) error {
	h.Logger.Info("Day 1")
	return nil
}
