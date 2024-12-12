package day1

import (
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Star1 is the solution for the first star
func Star1(h *common.Helpers) error {
	star := fmt.Sprintf("%d", star1)
	h.Logger.Info(human)
	f := h.Resources.GetFile(fmt.Sprintf("%s-%s", use, star))
	l, err := GetLists(h, f)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error getting lists: %s", err))
		return err
	}
	// print the lists
	h.Logger.Info(fmt.Sprintf("Left: %v", l.Left))
	h.Logger.Info(fmt.Sprintf("Right: %v", l.Right))

	return nil
}
