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
	h.Streams.Out.Write([]byte(fmt.Sprintf("%s\n", f.Name())))
	return nil
}
