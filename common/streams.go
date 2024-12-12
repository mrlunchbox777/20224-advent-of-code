package common

import (
	"io"
	"os"
)

// Streams is a struct that contains the input, output, and error streams
type Streams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer
}

// NewStreams creates a new Streams struct
func NewStreams(in io.Reader, out io.Writer, errOut io.Writer) *Streams {
	return &Streams{
		In:     in,
		Out:    out,
		ErrOut: errOut,
	}
}

// GetStreams returns a Streams struct with the default values
func GetStreams() *Streams {
	return NewStreams(os.Stdin, os.Stdout, os.Stderr)
}

// GetStreamsFunc is a function that returns a Streams struct with the default values
type GetStreamsFunc func() *Streams
