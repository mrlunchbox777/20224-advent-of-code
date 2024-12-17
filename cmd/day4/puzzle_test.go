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
		name      string
		input     *Block
		expected  *Block
		diagsInit bool
	}{
		{
			name: "rotate90_1x3",
			input: &Block{
				Rows: Sets{ms},
			},
			expected: &Block{
				Rows: Sets{
					{Cell{Letter: "M"}},
					{Cell{Letter: " "}},
					{Cell{Letter: "S"}},
				},
			},
			diagsInit: false,
		},
		{
			name: "rotate90_3x3",
			input: &Block{
				Rows: Sets{ms, a, ms}},
			expected: &Block{
				Rows: Sets{mm, a, ss}},
			diagsInit: true,
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
			assert.Empty(t, tc.input.RADiag)
			assert.Empty(t, tc.input.DDiag)
			assert.Empty(t, tc.input.RDDiag)
			if tc.diagsInit {
				assert.NotEmpty(t, result.ADiag)
				assert.NotEmpty(t, result.RADiag)
				assert.NotEmpty(t, result.DDiag)
				assert.NotEmpty(t, result.RDDiag)
			} else {
				assert.Empty(t, result.ADiag)
				assert.Empty(t, result.RADiag)
				assert.Empty(t, result.DDiag)
				assert.Empty(t, result.RDDiag)
			}
		})
	}
}

// TestRotate90x is a test for the rotate90x function
func TestRotate90x(t *testing.T) {
	testBlock_1x3 := &Block{
		Rows: Sets{ms},
	}
	testBlock_3x3 := &Block{
		Rows: Sets{
			ms,
			a,
			ms,
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
					{Cell{Letter: "M"}},
					{Cell{Letter: " "}},
					{Cell{Letter: "S"}},
				},
			},
		},
		{
			name:  "rotate90x_1x3_twice",
			input: testBlock_1x3,
			times: 2,
			expected: &Block{
				Rows: Sets{sm},
			},
		},
		{
			name:  "rotate90x_1x3_three_times",
			input: testBlock_1x3,
			times: 3,
			expected: &Block{
				Rows: Sets{
					{Cell{Letter: "S"}},
					{Cell{Letter: " "}},
					{Cell{Letter: "M"}},
				},
			},
		},
		{
			name:  "rotate90x_3x3_once",
			input: testBlock_3x3,
			times: 1,
			expected: &Block{
				Rows: Sets{
					mm,
					a,
					ss,
				},
			},
		},
		{
			name:  "rotate90x_3x3_twice",
			input: testBlock_3x3,
			times: 2,
			expected: &Block{
				Rows: Sets{
					sm,
					a,
					sm,
				},
			},
		},
		{
			name:  "rotate90x_3x3_three_times",
			input: testBlock_3x3,
			times: 3,
			expected: &Block{
				Rows: Sets{
					ss,
					a,
					mm,
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
