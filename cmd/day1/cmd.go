package day1

import (
	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

const (
	use   = "day1"
	human = "Day 1"
	star1 = 1
	star2 = 2
)

// NewCmd creates a new day1 command
func NewCmd(h *common.Helpers) *cobra.Command {
	day1Cmd := &cobra.Command{
		Use:   use,
		Short: human,
		Long:  human,
		RunE: func(cmd *cobra.Command, args []string) error {
			return day1(h)
		},
	}

	return day1Cmd
}

func day1(h *common.Helpers) error {
	if err := Star1(h); err != nil {
		return err
	}
	return nil
}
