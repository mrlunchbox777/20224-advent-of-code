package day2

import (
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

// Star2 is the solution for the first star
func Star2(h *common.Helpers) error {
	// l, err := getInputs(h, star1)
	// if err != nil {
	// 	h.Logger.Error(fmt.Sprintf("Error getting inputs: %s", err))
	// 	return err
	// }
	// _, err = h.Streams.Out.Write([]byte(fmt.Sprintf("%s Star 2: %d\n", human, l.CountCommonEntries(h))))
	// return err
	return nil
}
