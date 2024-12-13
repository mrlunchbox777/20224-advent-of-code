package day1

import (
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

// NewStar1Cmd creates a new star1 command
func NewStar1Cmd(h *common.Helpers) *cobra.Command {
	star1Cmd := &cobra.Command{
		Use:   star1,
		Short: star1,
		Long:  star1,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Star1(h)
		},
	}
	return star1Cmd
}

// Star1 is the solution for the first star
func Star1(h *common.Helpers) error {
	l, err := getInputs(h, star1)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error getting inputs: %s", err))
		return err
	}
	_, err = h.Streams.Out.Write([]byte(fmt.Sprintf("%s Star 1: %d\n", human, l.DiffList(h))))
	return err
}
