package day4

import (
	"testing"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/mrlunchbox777/2024-advent-of-code/common/test"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// TestRotate90 is a test for the rotate90 function
func TestRotate90(t *testing.T) {
	testCases := []struct {
		name     string
		input    *Block
		expected *Block
	}{
		{
			name: "rotate90",
			input: &Block{
				Rows: Sets{
					{
						{Letter: "M"},
						{Letter: " "},
						{Letter: "S"},
					},
				},
			},
			expected: &Block{
				Rows: Sets{
					{
						{Letter: "S"},
					},
					{
						{Letter: " "},
					},
					{
						{Letter: "M"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		// Arrange
		s := test.NewTestStreams()
		l := test.NewTestSlog(s.Streams)
		v := viper.New()
		h, err := common.NewHelpers(s.Streams, v, l)
		if err != nil {
			l.Error(err.Error())
			t.Log(err)
			t.Fail()
		}
		// Act
		result, err := tc.input.rotate90(h, false)
		// Assert
		assert.Nil(t, err)
		assert.Equal(t, result, tc.expected)
	}
}
