package test

import (
	"log/slog"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// NewTestSlog creates a new test slog
func NewTestSlog(s *common.Streams) *slog.Logger {
	o := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	h := slog.NewJSONHandler(s.ErrOut, o)
	l := slog.New(h)
	return l
}
