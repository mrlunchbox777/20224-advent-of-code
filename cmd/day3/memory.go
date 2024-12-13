package day3

import (
	"regexp"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Memory is a struct that contains the memory
type Memory struct {
	Raw             string
	RawGoodCommands []string
	GoodCommands    []Command
}

// Command is a struct that contains a command
type Command struct {
	Op   string
	Arg1 int
	Arg2 int
}

// GetMemory returns a new memory struct
func GetMemory(h *common.Helpers, in *common.File) (*Memory, error) {
	h.Logger.Debug("Getting memory")
	return parseInput(h, in)
}

// parseInput parses the input file and returns the memory
func parseInput(h *common.Helpers, in *common.File) (*Memory, error) {
	h.Logger.Debug("Parsing input")
	m := &Memory{
		Raw: string(in.Contents),
	}
	return m, nil
}

// findGoodCommands returns the good commands
func (m *Memory) findGoodCommands(h *common.Helpers) {
	h.Logger.Debug("Finding good commands")
	r := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)`)
	m.RawGoodCommands = r.FindAllString(m.Raw, -1)
}

// parseGoodCommands parses the good commands
func (m *Memory) parseGoodCommands(h *common.Helpers) {
	h.Logger.Debug("Parsing good commands")
	for _, c := range m.RawGoodCommands {
		r := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)`)
		matches := r.FindStringSubmatch(c)
		m.GoodCommands = append(m.GoodCommands, Command{
			Op:   matches[1],
			Arg1: h.ToInt(matches[2]),
			Arg2: h.ToInt(matches[3]),
		})
	}
}

// prepareMemory prepares the memory
func (m *Memory) prepareMemory(h *common.Helpers) {
	h.Logger.Debug("Preparing memory")
	m.findGoodCommands(h)
	m.parseGoodCommands(h)
}

// SumOfCommands returns the sum of the commands
func (m *Memory) SumOfCommands(h *common.Helpers) int {
	h.Logger.Debug("Summing commands")
	m.prepareMemory(h)
	sum := 0
	for _, c := range m.GoodCommands {
		sum += c.Arg1 * c.Arg2
	}
	return sum
}
