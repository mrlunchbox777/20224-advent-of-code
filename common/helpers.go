package common

import (
	"log/slog"

	"github.com/spf13/viper"
)

// Helpers is a struct that contains the common helpers
type Helpers struct {
	Streams *Streams
	Viper   *viper.Viper
	Logger  *slog.Logger
}

// NewHelpers creates a new Helpers struct
func NewHelpers(s *Streams, v *viper.Viper, l *slog.Logger) *Helpers {
	return &Helpers{
		Streams: s,
		Viper:   v,
		Logger:  l,
	}
}
