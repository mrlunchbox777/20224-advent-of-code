package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Lists has the left and right lists
type Lists struct {
	Left  []int
	Right []int
}

// GetLists returns the left and right lists
func GetLists(h *common.Helpers, in *common.File) (*Lists, error) {
	rawLists, err := parseInput(h, in)
	if err != nil {
		return nil, err
	}
	numLists := len(rawLists)
	if numLists > 2 {
		h.Logger.Error(fmt.Sprintf("Too many lists: %d", len(rawLists)))
		return nil, fmt.Errorf("Too many lists")
	}
	if numLists < 2 {
		h.Logger.Error(fmt.Sprintf("Not enough lists: %d", len(rawLists)))
		return nil, fmt.Errorf("Not enough lists")
	}
	leftLen := len(rawLists[0])
	rightLen := len(rawLists[1])
	if leftLen != rightLen {
		h.Logger.Error(fmt.Sprintf("Lists are not the same length: %d != %d", leftLen, rightLen))
		return nil, fmt.Errorf("Lists are not the same length")
	}
	lists := &Lists{
		Left:  rawLists[0],
		Right: rawLists[1],
	}
	return lists, nil
}

// parseInput parses the input file and returns the left and right lists
func parseInput(h *common.Helpers, in *common.File) ([][]int, error) {
	Left := []int{}
	Right := []int{}

	contents := string(in.Contents)
	h.Logger.Debug(fmt.Sprintf("Contents: %s", contents))
	lines := strings.Split(strings.ReplaceAll(contents, "\r\n", "\n"), "\n")
	h.Logger.Debug(fmt.Sprintf("Lines: %v", lines))
	h.Logger.Debug(fmt.Sprintf("Num lines: %d", len(lines)))

	for _, line := range lines {
		// Skip empty lines
		if line == "" {
			continue
		}
		// Split the line into words
		words := strings.Split(line, " ")
		added := 0
		for _, word := range words {
			// Skip empty words
			if word == "" {
				continue
			}
			// Convert the word to a number
			num, err := strconv.Atoi(word)
			if err != nil {
				return nil, err
			}
			// Add the number to the left or right list
			switch added {
			case 0:
				Left = append(Left, num)
				added++
			case 1:
				Right = append(Right, num)
				added++
			default:
				return nil, fmt.Errorf("Too many numbers in line")
			}
		}
		// Check if we have enough numbers
		if added != 2 {
			return nil, fmt.Errorf("Not enough numbers in line")
		}
	}
	h.Logger.Debug(fmt.Sprintf("Left: %v", Left))
	h.Logger.Debug(fmt.Sprintf("Right: %v", Right))

	return [][]int{Left, Right}, nil
}
