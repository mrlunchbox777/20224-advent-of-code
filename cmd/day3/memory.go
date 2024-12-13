package day3

import (
	"fmt"
	"regexp"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

const (
	// allCommandsMatch matches mul(1,2), do(), don't()
	allCommandsMatch = `((mul)\((\d{1,3}),(\d{1,3})\))|((do)\(\))|((don't)\(\))`
	// mulMatch matches mul(1,2)
	mulMatch = `(mul)\((\d{1,3}),(\d{1,3})\)`
	// doMatch matches do()
	doMatch = `(do)\(\)`
	// dontMatch matches don't()
	dontMatch = `(don't)\(\)`
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
	r := regexp.MustCompile(allCommandsMatch)
	m.RawGoodCommands = r.FindAllString(m.Raw, -1)
	h.Logger.Debug(fmt.Sprintf("Raw good commands: %v", m.RawGoodCommands))
}

// parseGoodCommands parses the good commands
func (m *Memory) parseGoodCommands(h *common.Helpers) {
	h.Logger.Debug("Parsing good commands")
	for _, c := range m.RawGoodCommands {
		rm := regexp.MustCompile(mulMatch)
		matches := rm.FindStringSubmatch(c)
		if len(matches) > 0 {
			m.GoodCommands = append(m.GoodCommands, Command{
				Op:   matches[1],
				Arg1: h.ToInt(matches[2]),
				Arg2: h.ToInt(matches[3]),
			})
			continue
		}
		rd := regexp.MustCompile(doMatch)
		matches = rd.FindStringSubmatch(c)
		if len(matches) > 0 {
			m.GoodCommands = append(m.GoodCommands, Command{
				Op: matches[1],
			})
			continue
		}
		rdn := regexp.MustCompile(dontMatch)
		matches = rdn.FindStringSubmatch(c)
		if len(matches) > 0 {
			m.GoodCommands = append(m.GoodCommands, Command{
				Op: matches[1],
			})
			continue
		}
	}
	h.Logger.Debug(fmt.Sprintf("Good commands: %v", m.GoodCommands))
}

// prepareMemory prepares the memory
func (m *Memory) prepareMemory(h *common.Helpers) {
	h.Logger.Debug("Preparing memory")
	m.findGoodCommands(h)
	m.parseGoodCommands(h)
}

// SumOfCommands returns the sum of the commands
func (m *Memory) SumOfCommands(h *common.Helpers, flowControl bool) int {
	h.Logger.Debug("Summing commands")
	m.prepareMemory(h)
	sum := 0
	enabled := true
	for _, c := range m.GoodCommands {
		if flowControl {
			if c.Op == "do" {
				enabled = true
			}
			if c.Op == "don't" {
				enabled = false
			}
		}
		if enabled && c.Op == "mul" {
			sum += c.Arg1 * c.Arg2
		}
	}
	return sum
}
