package day4

import (
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

const (
	use   = "day4"
	human = "Day 4"
	star1 = "star1"
	star2 = "star2"
)

// NewCmd creates a new day1 command
func NewCmd(h *common.Helpers) *cobra.Command {
	day1Cmd := &cobra.Command{
		Use:   use,
		Short: human,
		Long:  human,
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("No subcommand given")
		},
	}

	day1Cmd.AddCommand(NewStar1Cmd(h))
	day1Cmd.AddCommand(NewStar2Cmd(h))

	return day1Cmd
}

func getInputs(h *common.Helpers, star string) (*Puzzle, error) {
	resourceName := fmt.Sprintf("%s-%s", use, star1)
	name := fmt.Sprintf("%s-%s", use, star)
	h.Logger.Info(name)
	f := h.Resources.GetFile(h, resourceName)
	p, err := GetPuzzle(h, f)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error getting puzzle: %s", err))
		return nil, err
	}
	// print the lists
	h.Logger.Debug(fmt.Sprintf("Puzzle: %v", p))
	return p, nil
}
