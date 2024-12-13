package day4

import (
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

// Puzzle is the struct for a word search puzzle
type Puzzle struct {
	Raw string
	Block
}

type Block struct {
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
		Raw: string(in.Contents),
		Block: Block{
			Initialized: false,
		},
	}
	p.getRows(h)
	return p, nil
}

// getRows parses the rows, doesn't initialize the puzzle
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
}

// getCols parses the columns
func (p *Block) getCols(h *common.Helpers) {
	h.Logger.Debug("Getting columns")
	p.Cols = make([]Set, len(p.Rows[0]))
	for _, r := range p.Rows {
		for i, c := range r {
			p.Cols[i] = append(p.Cols[i], c)
		}
	}
	p.RCols = p.getRSets(h, p.Cols)
}

// getADiag parses the ascending diagonals, starting from the top left
func (p *Block) getADiag(h *common.Helpers) {
	h.Logger.Debug("Getting ascending diagonals")
	p.ADiag = make([]Set, len(p.Rows)*2-1)
	for y := 0; y < len(p.Rows); y++ {
		for x := 0; x < len(p.Rows[y]); x++ {
			p.ADiag[y+x] = append(p.ADiag[y+x], p.Rows[y][x])
		}
	}
	p.RADiag = p.getRSets(h, p.ADiag)
}

// getDDiag parses the descending diagonals, starting from the top right
func (p *Block) getDDiag(h *common.Helpers) {
	h.Logger.Debug("Getting descending diagonals")
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
func (p *Block) getRSets(h *common.Helpers, s Sets) Sets {
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

// initialize initializes the block
func (b *Block) initialize(h *common.Helpers, rows Sets) error {
	if rows == nil || len(rows) == 0 {
		h.Logger.Error("No rows to initialize")
		return fmt.Errorf("No rows to initialize")
	}
	h.Logger.Debug("Initializing block: %i x %i", len(rows), len(rows[0]))
	b.Rows = rows
	b.RRows = b.getRSets(h, b.Rows)
	b.getCols(h)
	b.getADiag(h)
	b.getDDiag(h)
	b.Initialized = true
	return nil
}

// CountWord returns the number of times a word appears in the puzzle
func (p *Puzzle) CountWord(h *common.Helpers, word string) (int, error) {
	return p.countWordInBlock(h, word)
}

// countWordInSets returns the number of times a word appears in a set of cells
func (p *Block) countWordInSets(sets Sets, word string) int {
	count := 0
	for _, s := range sets {
		count += p.countWordInSet(s, word)
	}
	return count
}

// countWordInSet returns the number of times a word appears in a set of cells
func (p *Block) countWordInSet(set Set, word string) int {
	count := 0
	for i := 0; i < len(set)-len(word)+1; i++ {
		if p.checkSegment(set[i:i+len(word)], word) {
			count++
		}
	}
	return count
}

// checkSegment checks if a segment contains a word
func (p *Block) checkSegment(cells []Cell, word string) bool {
	for i, c := range cells {
		if c.Letter != string(word[i]) {
			return false
		}
	}
	return true
}

// getBlock returns the block for a given Sets
func (p *Puzzle) getBlock(h *common.Helpers, sets Sets) (*Block, error) {
	h.Logger.Debug("Getting block")
	b := &Block{
		Initialized: false,
	}
	err := b.initialize(h, sets)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Error initializing block: %s", err))
		return nil, err
	}
	return b, nil
}

// countWordInBlock returns the number of times a word appears in a block
func (b *Block) countWordInBlock(h *common.Helpers, word string) (int, error) {
	h.Logger.Debug("Counting word in block")
	if !b.Initialized {
		err := b.initialize(h, b.Rows)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error initializing block: %s", err))
			return 0, err
		}
	}
	count := 0
	count += b.countWordInSets(b.Rows, word)
	count += b.countWordInSets(b.RRows, word)
	count += b.countWordInSets(b.Cols, word)
	count += b.countWordInSets(b.RCols, word)
	count += b.countWordInSets(b.DDiag, word)
	count += b.countWordInSets(b.RDDiag, word)
	count += b.countWordInSets(b.ADiag, word)
	count += b.countWordInSets(b.RADiag, word)
	return count, nil
}
