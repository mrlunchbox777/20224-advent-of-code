package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Level is an int that represents the level of the report
type Level int

// Report is a slice that contains the report
type Report []Level

// Reports is a slice of reports
type Reports []Report

// GetReports returns a slice of reports
func GetReports(h *common.Helpers, in *common.File) (*Reports, error) {
	h.Logger.Debug("Getting reports")
	return parseInput(h, in)
}

// parseInput parses the input file and returns the reports
func parseInput(h *common.Helpers, in *common.File) (*Reports, error) {
	h.Logger.Debug("Parsing input")
	reports := &Reports{}
	contents := string(in.Contents)
	lines := h.GetLines(contents)
	for _, l := range lines {
		// Skip empty lines
		if l == "" {
			continue
		}
		rawLevels := strings.Split(l, " ")
		r := Report{}
		for _, rl := range rawLevels {
			// Skip empty levels
			if rl == "" {
				continue
			}
			// parse the level
			intLevel, err := strconv.Atoi(rl)
			if err != nil {
				h.Logger.Error(fmt.Sprintf("Error parsing level: %s", err))
				return nil, err
			}
			r = append(r, Level(intLevel))
		}
		*reports = append(*reports, r)
	}
	return reports, nil
}

// checkDampener checks if the dampener can handle the error
func checkDampener(h *common.Helpers, dampener bool, dampened bool) bool {
	h.Logger.Debug("Checking dampener")
	if dampener && !dampened {
		h.Logger.Debug("Dampening")
		dampened = true
		return true
	}
	return false
}

// IsSafe returns true if the report is safe
func (r *Report) IsSafe(h *common.Helpers, dampener bool) bool {
	h.Logger.Debug("Checking if report is safe")
	count := 0
	increasing := false
	prev := Level(0)
	dampened := false
	for _, l := range *r {
		if count > 0 {
			if count == 1 || (count == 2 && dampened) {
				increasing = l > prev
			}
			// must stay increasing
			if increasing && l <= prev {
				if checkDampener(h, dampener, dampened) {
					dampened = true
					continue
				}
				return false
			}
			// must stay decreasing
			if !increasing && l >= prev {
				if checkDampener(h, dampener, dampened) {
					dampened = true
					continue
				}
				return false
			}
			// must not decrease by more than 3
			if l < prev && l < (prev-3) {
				if checkDampener(h, dampener, dampened) {
					dampened = true
					continue
				}
				return false
			}
			// must not increase by more than 3
			if l > prev && l > (prev+3) {
				if checkDampener(h, dampener, dampened) {
					dampened = true
					continue
				}
				return false
			}
		}
		prev = l
		count++
	}
	return true
}

// CountSafeEntries returns the number of safe reports
func (r *Reports) CountSafeEntries(h *common.Helpers, dampener bool) int {
	count := 0
	for _, report := range *r {
		if report.IsSafe(h, dampener) {
			count++
		}
	}
	return count
}
