package day4

import (
	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Puzzle is the struct for a word search puzzle
type Puzzle struct {
	Raw         string
	Initialized bool
	Rows        Sets
	RRows       Sets
	Cols        Sets
	RCols       Sets
	DDiag       Sets
	RDDiag      Sets
	ADiag       Sets
	RADiag      Sets
}

// Sets is a slice of sets
type Sets []Set

// Set is the struct for a row/col/diag in a word search puzzle
type Set []Cell

// Cell is the struct for a cell in a word search puzzle
type Cell struct {
	Letter string
}

// GetPuzzle returns a new puzzle struct
func GetPuzzle(h *common.Helpers, in *common.File) (*Puzzle, error) {
	h.Logger.Debug("Getting puzzle")
	return parseInput(h, in)
}

// parseInput parses the input file and returns the puzzle
func parseInput(h *common.Helpers, in *common.File) (*Puzzle, error) {
	h.Logger.Debug("Parsing input")
	p := &Puzzle{
		Raw:         string(in.Contents),
		Initialized: false,
		Rows:        []Set{},
		RRows:       []Set{},
		Cols:        []Set{},
		RCols:       []Set{},
		DDiag:       []Set{},
		RDDiag:      []Set{},
		ADiag:       []Set{},
		RADiag:      []Set{},
	}
	return p, nil
}

// getRows parses the rows
func (p *Puzzle) getRows(h *common.Helpers) {
	h.Logger.Debug("Getting rows")
	lines := h.GetLines(p.Raw)
	for _, line := range lines {
		var row Set
		for _, letter := range line {
			row = append(row, Cell{Letter: string(letter)})
		}
		p.Rows = append(p.Rows, row)
	}
	p.RRows = p.getRSets(h, p.Rows)
}

// getCols parses the columns
func (p *Puzzle) getCols(h *common.Helpers) {
	h.Logger.Debug("Getting columns")
	if len(p.Rows) == 0 {
		p.getRows(h)
	}
	p.Cols = make([]Set, len(p.Rows[0]))
	for _, r := range p.Rows {
		for i, c := range r {
			p.Cols[i] = append(p.Cols[i], c)
		}
	}
	p.RCols = p.getRSets(h, p.Cols)
}

// getADiag parses the ascending diagonals, starting from the top left
func (p *Puzzle) getADiag(h *common.Helpers) {
	h.Logger.Debug("Getting ascending diagonals")
	if len(p.Rows) == 0 {
		p.getRows(h)
	}
	p.ADiag = make([]Set, len(p.Rows)*2-1)
	for y := 0; y < len(p.Rows); y++ {
		for x := 0; x < len(p.Rows[y]); x++ {
			p.ADiag[y+x] = append(p.ADiag[y+x], p.Rows[y][x])
		}
	}
	p.RADiag = p.getRSets(h, p.ADiag)
}

// getDDiag parses the descending diagonals, starting from the top right
func (p *Puzzle) getDDiag(h *common.Helpers) {
	h.Logger.Debug("Getting descending diagonals")
	if len(p.Rows) == 0 {
		p.getRows(h)
	}
	p.DDiag = make([]Set, len(p.Rows)*2-1)
	for y := 0; y < len(p.Rows); y++ {
		xLen := len(p.Rows[y]) - 1
		for x := 0; x < len(p.Rows[y]); x++ {
			p.DDiag[y+x] = append(p.DDiag[y+x], p.Rows[y][xLen-x])
		}
	}
	p.RDDiag = p.getRSets(h, p.DDiag)
}

// getRSets the reverse set
func (p *Puzzle) getRSets(h *common.Helpers, s Sets) Sets {
	h.Logger.Debug("Getting reverse set")
	var rs Sets
	for _, r := range s {
		var rSet Set
		for i := len(r) - 1; i >= 0; i-- {
			rSet = append(rSet, r[i])
		}
		rs = append(rs, rSet)
	}
	return rs
}

// initialize initializes the puzzle
func (p *Puzzle) initialize(h *common.Helpers) {
	h.Logger.Debug("Initializing puzzle")
	p.getRows(h)
	p.getCols(h)
	p.getADiag(h)
	p.getDDiag(h)
	p.Initialized = true
}

// CountWord returns the number of times a word appears in the puzzle
func (p *Puzzle) CountWord(h *common.Helpers, word string) int {
	h.Logger.Debug("Counting word")
	if !p.Initialized {
		p.initialize(h)
	}
	count := 0
	count += p.countWordInSets(p.Rows, word)
	count += p.countWordInSets(p.RRows, word)
	count += p.countWordInSets(p.Cols, word)
	count += p.countWordInSets(p.RCols, word)
	count += p.countWordInSets(p.DDiag, word)
	count += p.countWordInSets(p.RDDiag, word)
	count += p.countWordInSets(p.ADiag, word)
	count += p.countWordInSets(p.RADiag, word)
	return count
}

// countWordInSets returns the number of times a word appears in a set of cells
func (p *Puzzle) countWordInSets(sets Sets, word string) int {
	count := 0
	for _, s := range sets {
		count += p.countWordInSet(s, word)
	}
	return count
}

// countWordInSet returns the number of times a word appears in a set of cells
func (p *Puzzle) countWordInSet(set Set, word string) int {
	count := 0
	for i := 0; i < len(set)-len(word)+1; i++ {
		if p.checkSegment(set[i:i+len(word)], word) {
			count++
		}
	}
	return count
}

// checkSegment checks if a segment contains a word
func (p *Puzzle) checkSegment(cells []Cell, word string) bool {
	for i, c := range cells {
		if c.Letter != string(word[i]) {
			return false
		}
	}
	return true
}
