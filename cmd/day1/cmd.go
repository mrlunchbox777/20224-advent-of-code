package day1

import (
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

const (
	use   = "day1"
	human = "Day 1"
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

func getInputs(h *common.Helpers, star string) (*Lists, error) {
	resourceName := fmt.Sprintf("%s-%s", use, star1)
	name := fmt.Sprintf("%s-%s", use, star)
	h.Logger.Info(name)
	f := h.Resources.GetFile(h, resourceName)
	l, err := GetLists(h, f)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error getting lists: %s", err))
		return nil, err
	}
	l.Sort(h)
	// print the lists
	h.Logger.Debug(fmt.Sprintf("Left: %v", l.Left))
	h.Logger.Debug(fmt.Sprintf("Right: %v", l.Right))
	return l, nil
}
