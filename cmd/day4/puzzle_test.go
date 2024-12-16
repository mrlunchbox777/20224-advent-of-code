package day4

import (
	"testing"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
	"github.com/mrlunchbox777/2024-advent-of-code/common/test"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	ms = Set{
		Cell{Letter: "M"},
		Cell{Letter: " "},
		Cell{Letter: "S"},
	}
	a = Set{
		Cell{Letter: " "},
		Cell{Letter: "A"},
		Cell{Letter: " "},
	}
	sm = Set{
		Cell{Letter: "S"},
		Cell{Letter: " "},
		Cell{Letter: "M"},
	}
	mm = Set{
		Cell{Letter: "M"},
		Cell{Letter: " "},
		Cell{Letter: "M"},
	}
	ss = Set{
		Cell{Letter: "S"},
		Cell{Letter: " "},
		Cell{Letter: "S"},
	}
)

// TestRotate90 is a test for the rotate90 function
func TestRotate90(t *testing.T) {
	testCases := []struct {
		name     string
		input    *Block
		expected *Block
	}{
		{
			name: "rotate90_1x3",
			input: &Block{
				Rows: Sets{ms},
			},
			expected: &Block{
				Rows: Sets{sm},
			},
		},
		{
			name: "rotate90_3x3",
			input: &Block{
				Rows: Sets{ms, a, ms}},
			expected: &Block{
				Rows: Sets{mm, a, ss}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
			assert.Equal(t, tc.expected, result)
			// Act
			err = result.initialize(h, result.Rows)
			// Assert
			assert.Nil(t, err)
			assert.Equal(t, tc.expected.Rows, result.Rows)
			assert.Empty(t, tc.input.RRows)
			assert.NotEqual(t, tc.input.RRows, result.RRows)
			assert.Empty(t, tc.input.Cols)
			assert.NotEqual(t, tc.input.Cols, result.Cols)
			assert.Empty(t, tc.input.RCols)
			assert.NotEqual(t, tc.input.RCols, result.RCols)
			assert.Empty(t, tc.input.ADiag)
			assert.NotEqual(t, tc.input.ADiag, result.ADiag)
			assert.Empty(t, tc.input.RADiag)
			assert.NotEqual(t, tc.input.RADiag, result.RADiag)
			assert.Empty(t, tc.input.DDiag)
			assert.NotEqual(t, tc.input.DDiag, result.DDiag)
			assert.Empty(t, tc.input.RDDiag)
			assert.NotEqual(t, tc.input.RDDiag, result.RDDiag)
		})
	}
}

// TestRotate90x is a test for the rotate90x function
func TestRotate90x(t *testing.T) {
	testBlock_1x3 := &Block{
		Rows: Sets{
			{
				{Letter: "M"},
				{Letter: " "},
				{Letter: "S"},
			},
		},
	}
	testBlock_3x3 := &Block{
		Rows: Sets{
			{
				{Letter: "M"},
				{Letter: " "},
				{Letter: "S"},
			},
			{
				{Letter: " "},
				{Letter: "A"},
				{Letter: " "},
			},
			{
				{Letter: "M"},
				{Letter: " "},
				{Letter: "S"},
			},
		},
	}
	testCases := []struct {
		name     string
		input    *Block
		times    int
		expected *Block
	}{
		{
			name:  "rotate90x_1x3_once",
			input: testBlock_1x3,
			times: 1,
			expected: &Block{
				Rows: Sets{
					{
						{Letter: "M"},
					},
					{
						{Letter: " "},
					},
					{
						{Letter: "S"},
					},
				},
			},
		},
		{
			name:  "rotate90x_1x3_twice",
			input: testBlock_1x3,
			times: 2,
			expected: &Block{
				Rows: Sets{
					{
						{Letter: "S"},
						{Letter: " "},
						{Letter: "M"},
					},
				},
			},
		},
		{
			name:  "rotate90x_1x3_three_times",
			input: testBlock_1x3,
			times: 3,
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
		{
			name:  "rotate90x_3x3_once",
			input: testBlock_3x3,
			times: 1,
			expected: &Block{
				Rows: Sets{
					{
						{Letter: "M"},
						{Letter: " "},
						{Letter: "M"},
					},
					{
						{Letter: " "},
						{Letter: "A"},
						{Letter: " "},
					},
					{
						{Letter: "S"},
						{Letter: " "},
						{Letter: "S"},
					},
				},
			},
		},
		{
			name:  "rotate90x_3x3_twice",
			input: testBlock_3x3,
			times: 2,
			expected: &Block{
				Rows: Sets{
					{
						{Letter: "S"},
						{Letter: " "},
						{Letter: "M"},
					},
					{
						{Letter: " "},
						{Letter: "A"},
						{Letter: " "},
					},
					{
						{Letter: "M"},
						{Letter: " "},
						{Letter: "S"},
					},
				},
			},
		},
		{
			name:  "rotate90x_3x3_three_times",
			input: testBlock_3x3,
			times: 3,
			expected: &Block{
				Rows: Sets{
					{
						{Letter: "S"},
						{Letter: " "},
						{Letter: "S"},
					},
					{
						{Letter: " "},
						{Letter: "A"},
						{Letter: " "},
					},
					{
						{Letter: "M"},
						{Letter: " "},
						{Letter: "M"},
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
			result, err := tc.input.rotate90x(h, tc.times)
			// Assert
			assert.Nil(t, err)
			assert.Equal(t, tc.expected.Rows, result.Rows)
		})
	}
}
