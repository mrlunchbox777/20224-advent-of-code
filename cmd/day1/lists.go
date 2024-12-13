package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Lists has the left and right lists
type Lists struct {
	Left        []int
	Right       []int
	LeftCounts  map[int]int
	RightCounts map[int]int
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
		Left:        rawLists[0],
		Right:       rawLists[1],
		LeftCounts:  make(map[int]int),
		RightCounts: make(map[int]int),
	}
	lists.indexInstances(h)
	return lists, nil
}

// parseInput parses the input file and returns the left and right lists
func parseInput(h *common.Helpers, in *common.File) ([][]int, error) {
	Left := []int{}
	Right := []int{}

	h.Logger.Debug(fmt.Sprintf("Parsing input: %s", in.Name))
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

// Sort sorts the lists
func (l *Lists) Sort(h *common.Helpers) {
	h.Logger.Debug("Sorting lists")
	sortList(h, l.Left)
	sortList(h, l.Right)
}

// sortList sorts a list
func sortList(h *common.Helpers, l []int) {
	h.Logger.Debug(fmt.Sprintf("Sorting list: %v", l))
	slices.Sort(l)
}

// diffListEntry returns the difference between the left and right lists at index i
func diffListEntry(h *common.Helpers, l *Lists, i int) int {
	h.Logger.Debug(fmt.Sprintf("DiffListEntry: %v", l))
	if l.Left[i] < l.Right[i] {
		return l.Right[i] - l.Left[i]
	}
	return l.Left[i] - l.Right[i]
}

// DiffList returns the difference between the left and right lists
func (l *Lists) DiffList(h *common.Helpers) int {
	h.Logger.Debug(fmt.Sprintf("DiffList: %v", l))
	diff := 0
	for i := 0; i < len(l.Left); i++ {
		diff += diffListEntry(h, l, i)
	}
	return diff
}

// indexInstances returns the number of instances of a number in a list
func (l *Lists) indexInstances(h *common.Helpers) {
	h.Logger.Debug(fmt.Sprintf("IndexInstances: %v", l))
	for _, num := range l.Left {
		l.LeftCounts[num]++
	}
	for _, num := range l.Right {
		l.RightCounts[num]++
	}
}

// weightOfIndex returns the weight of an index
func (l *Lists) weightOfIndex(h *common.Helpers, i int) int {
	h.Logger.Debug(fmt.Sprintf("WeightOfIndex: %v", i))
	leftValue := l.Left[i]
	rightCount := l.RightCounts[leftValue]
	return leftValue * rightCount
}

// CountCommonEntries returns the product of the number of a common entry in the left and right lists
func (l *Lists) CountCommonEntries(h *common.Helpers) int {
	h.Logger.Debug(fmt.Sprintf("CountCommonEntries: %v", l))
	total := 0
	for i, _ := range l.Left {
		total += l.weightOfIndex(h, i)
	}
	return total
}
