package test

import (
	"bytes"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// NewTestStreams creates a new TestStreams struct
func NewTestStreams() *TestStreams {
	In := bytes.NewBufferString("")
	Out := bytes.NewBufferString("")
	ErrOut := bytes.NewBufferString("")
	return &TestStreams{
		BufIn:       In,
		BufInOut:    Out,
		BufInErrOut: ErrOut,
		Streams: &common.Streams{
			In:     In,
			Out:    Out,
			ErrOut: ErrOut,
		},
	}
}

// TestStreams is a struct that contains the input, output, and error streams
type TestStreams struct {
	*common.Streams
	BufIn       *bytes.Buffer
	BufInOut    *bytes.Buffer
	BufInErrOut *bytes.Buffer
}
