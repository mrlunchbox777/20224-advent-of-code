package common

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

// Helpers is a struct that contains the common helpers
type Helpers struct {
	Streams   *Streams
	Viper     *viper.Viper
	Logger    *slog.Logger
	Resources *Resources
}

// ErrStreamsNil is an error that is returned when the streams is nil
type ErrStreamsNil struct{}

// Error returns the error message
func (e ErrStreamsNil) Error() string {
	return "streams is nil"
}

// ErrViperNil is an error that is returned when the viper is nil
type ErrViperNil struct{}

// Error returns the error message
func (e ErrViperNil) Error() string {
	return "viper is nil"
}

// ErrLoggerNil is an error that is returned when the logger is nil
type ErrLoggerNil struct{}

// Error returns the error message
func (e ErrLoggerNil) Error() string {
	return "logger is nil"
}

// NewHelpers creates a new Helpers struct
func NewHelpers(s *Streams, v *viper.Viper, l *slog.Logger) (*Helpers, error) {
	if l == nil {
		return nil, ErrLoggerNil{}
	}
	var err error
	if s == nil {
		err = ErrStreamsNil{}
	}
	if v == nil {
		err = ErrViperNil{}
	}
	if err != nil {
		l.Error(fmt.Sprintf("helpers: %s", err.Error()))
		return nil, err
	}
	r, err := NewResources(l, v)
	if err != nil {
		return nil, err
	}
	return &Helpers{
		Streams:   s,
		Viper:     v,
		Logger:    l,
		Resources: r,
	}, nil
}

// GetLines returns a slice of lines from a string
func (h *Helpers) GetLines(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}
