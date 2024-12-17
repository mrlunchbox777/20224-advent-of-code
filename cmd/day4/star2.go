package day4

import (
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/spf13/cobra"
)

// NewStar2Cmd creates a new star2 command
func NewStar2Cmd(h *common.Helpers) *cobra.Command {
	star2Cmd := &cobra.Command{
		Use:   star2,
		Short: star2,
		Long:  star2,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Star2(h)
		},
	}
	return star2Cmd
}

// Star2 is the solution for the second star
func Star2(h *common.Helpers) error {
	p, err := getInputs(h, star2)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error getting inputs: %s", err))
		return err
	}
	// create sub sets
	ms := Set{
		Cell{Letter: "M"},
		Cell{Letter: " "},
		Cell{Letter: "S"},
	}
	a := Set{
		Cell{Letter: " "},
		Cell{Letter: "A"},
		Cell{Letter: " "},
	}
	// create the target sets
	sets := []Sets{
		{
			ms,
			a,
			ms,
		},
	}
	count, err := p.CountBlocks(h, sets, true)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error counting word: %s", err))
		return err
	}
	_, err = h.Streams.Out.Write([]byte(fmt.Sprintf("%s Star 2: %d\n", human, count)))
	return err
}
