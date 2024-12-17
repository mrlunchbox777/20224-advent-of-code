package day4

import (
	"encoding/json"
	"fmt"

	"github.com/mrlunchbox777/2024-advent-of-code/common"
)

const (
	// Column is the column size type
	Column = WrongSizeType("Column")
	// Row is the row size type
	Row = WrongSizeType("Row")
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
	Size        *Size
}

// Sets is a slice of sets
type Sets []Set

// Set is the struct for a row/col/diag in a word search puzzle
type Set []Cell

// Cell is the struct for a cell in a word search puzzle
type Cell struct {
	Letter string
}

// Size is the struct for the size of a block
type Size struct {
	X int
	Y int
}

// Equals checks if two sizes are equal
func (s *Size) Equals(o *Size) bool {
	return s.X == o.X && s.Y == o.Y
}

// SubBlocks is a slice of blocks by size
type SubBlocks map[*Size][]*Block

// WrongSizeType is an int that represents the type of wrong size error
type WrongSizeType string

// WrongSizeError is an error for when the sizes don't match
type WrongSizeError struct {
	Expected int
	Actual   int
	Type     WrongSizeType
}

// Error returns the error message
func (e *WrongSizeError) Error() string {
	return fmt.Sprintf("Wrong %s size. Expected %d, got %d", e.Type, e.Expected, e.Actual)
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

// getBlock returns the block for a given Sets
func getBlock(h *common.Helpers, sets Sets) (*Block, error) {
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
	p.Cols = make([]Set, p.Size.X)
	for _, r := range p.Rows {
		for i, c := range r {
			p.Cols[i] = append(p.Cols[i], c)
		}
	}
	p.RCols = p.getRSets(h, p.Cols)
}

// getADiag parses the ascending diagonals, starting from the top left
func (p *Block) getADiag(h *common.Helpers) {
	if p.Size.Y != p.Size.X {
		h.Logger.Debug("Can't get ascending diagonals")
		return
	}
	h.Logger.Debug("Getting ascending diagonals")
	p.ADiag = make([]Set, p.Size.Y*2-1)
	for y := 0; y < p.Size.Y; y++ {
		for x := 0; x < p.Size.X; x++ {
			p.ADiag[y+x] = append(p.ADiag[y+x], p.Rows[y][x])
		}
	}
	p.RADiag = p.getRSets(h, p.ADiag)
}

// getDDiag parses the descending diagonals, starting from the top right
func (p *Block) getDDiag(h *common.Helpers) {
	if p.Size.Y != p.Size.X {
		h.Logger.Debug("Can't get descending diagonals")
		return
	}
	h.Logger.Debug("Getting descending diagonals")
	p.DDiag = make([]Set, p.Size.Y*2-1)
	for y := 0; y < p.Size.Y; y++ {
		xLen := p.Size.X - 1
		for x := 0; x < p.Size.X; x++ {
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
	if b.Initialized {
		h.Logger.Debug("Block already initialized")
		return nil
	}
	if len(rows) == 0 {
		h.Logger.Error("No rows to initialize")
		return fmt.Errorf("No rows to initialize")
	}
	h.Logger.Debug(fmt.Sprintf("Initializing block: %d x %d", len(rows), len(rows[0])))
	b.Rows = rows
	b.Size = &Size{
		X: len(b.Rows[0]),
		Y: len(b.Rows),
	}
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

// getBlocksFromBlock returns the subblocks from a block based on a target size
func (b *Block) getBlocksFromBlock(h *common.Helpers, target *Size) ([]*Block, error) {
	h.Logger.Debug("Getting blocks from block")
	h.Logger.Debug(fmt.Sprintf("Target: %d x %d", target.X, target.Y))
	h.Logger.Debug(fmt.Sprintf("Block: %d x %d", b.Size.X, b.Size.Y))
	blocks := make([]*Block, 0)
	// get dimensions of target
	// get the first row of target sized blocks from the source
	for y := 0; y < b.Size.Y-target.Y; y++ {
		for x := 0; x < b.Size.X-target.X; x++ {
			block := &Block{}
			// get the subset of rows
			for i := 0; i < target.Y; i++ {
				row := make(Set, 0)
				// only get the subset of columns
				row = append(row, b.Rows[y+i][x:x+target.X]...)
				block.Rows = append(block.Rows, row)
			}
			h.Logger.Debug(fmt.Sprintf("Block: %d x %d", target.X, target.Y))
			err := block.initialize(h, block.Rows)
			if err != nil {
				h.Logger.Error(fmt.Sprintf("Error initializing block: %s", err))
				return nil, err
			}
			blocks = append(blocks, block)
		}
	}
	return blocks, nil
}

// doBlocksMatch checks if two blocks match, use " " for wildcards
func (b *Block) doBlocksMatch(h *common.Helpers, target *Block) (bool, error) {
	h.Logger.Debug("Checking if blocks match")
	if b.Size.Y != target.Size.Y {
		return false, &WrongSizeError{
			Expected: b.Size.Y,
			Actual:   target.Size.Y,
			Type:     Row,
		}
	}
	if b.Size.X != target.Size.X {
		return false, &WrongSizeError{
			Expected: b.Size.X,
			Actual:   target.Size.X,
			Type:     Column,
		}
	}
	// every cell must match
	for y := 0; y < b.Size.Y; y++ {
		for x := 0; x < b.Size.X; x++ {
			// skip wildcards
			if target.Rows[y][x].Letter == " " {
				continue
			}
			if b.Rows[y][x].Letter != target.Rows[y][x].Letter {
				return false, nil
			}
		}
	}
	return true, nil
}

// rotate90 rotates the block 90 degrees
func (b *Block) rotate90(h *common.Helpers, init bool) (*Block, error) {
	h.Logger.Debug("Rotating block")
	if b.Size.Y == 0 {
		h.Logger.Error("No rows to rotate")
		return nil, fmt.Errorf("No rows to rotate")
	}
	block := &Block{}
	block.Rows = make(Sets, b.Size.X)
	for x := 0; x < b.Size.X; x++ {
		block.Rows[x] = make(Set, b.Size.Y)
		for y := 0; y < b.Size.Y; y++ {
			block.Rows[x][y] = b.Rows[b.Size.Y-y-1][x]
		}
	}
	if init {
		err := block.initialize(h, block.Rows)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error initializing block: %s", err))
			return nil, err
		}
	} else {
		block.Size = &Size{
			X: b.Size.Y,
			Y: b.Size.X,
		}
	}
	return block, nil
}

// rotate90x rotates the block 90 degrees
func (b *Block) rotate90x(h *common.Helpers, count int) (*Block, error) {
	h.Logger.Debug("Rotating block")
	block := b
	var err error
	for i := 0; i < count; i++ {
		block, err = block.rotate90(h, false)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error rotating block: %s", err))
			return nil, err
		}
	}
	err = block.initialize(h, block.Rows)
	return block, err
}

// doBlocksMatchAny checks if two blocks match in any orientation (rotated and checked 3 times), use " " for wildcards
func (b *Block) doBlocksMatchAny(h *common.Helpers, targets []*Block) (int, error) {
	h.Logger.Debug("Checking if blocks match")
	count := 0
	for _, target := range targets {
		match, err := b.doBlocksMatch(h, target)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error checking if blocks match: %s", err))
			return 0, err
		}
		if match {
			count++
		}
	}
	return count, nil
}

// countBlockInBlock returns the number of times a Block appears in a block (use " " for wildcards)
func (b *Block) countBlockInBlockSameSize(h *common.Helpers, targets []Sets, rotate bool) (int, error) {
	h.Logger.Debug("Counting block in block")
	if !b.Initialized {
		err := b.initialize(h, b.Rows)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error initializing block: %s", err))
			return 0, err
		}
	}
	targetBlocks := make([]*Block, 0)
	for _, target := range targets {
		targetBlock, err := getBlock(h, target)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error getting block: %s", err))
			return 0, err
		}
		targetBlocks = append(targetBlocks, targetBlock)
	}
	if rotate {
		for _, target := range targets {
			targetBlock, err := getBlock(h, target)
			if err != nil {
				h.Logger.Error(fmt.Sprintf("Error getting block: %s", err))
				return 0, err
			}
			// already have the original block
			for i := 1; i < 4; i++ {
				newTarget, err := targetBlock.rotate90x(h, i)
				if err != nil {
					h.Logger.Error(fmt.Sprintf("Error rotating block: %s", err))
					return 0, err
				}
				h.Logger.Debug(fmt.Sprintf("Rotated block: %d x %d", newTarget.Size.X, newTarget.Size.Y))
				stringNewTarget, err := json.Marshal(newTarget)
				if err != nil {
					h.Logger.Error(fmt.Sprintf("Error marshalling block: %s", err))
					return 0, err
				}
				h.Logger.Debug(fmt.Sprintf("Rotated block values: %s", stringNewTarget))
				targetBlocks = append(targetBlocks, newTarget)
			}
		}
	}

	targetSizes := make([]*Size, 0)
	for _, targetBlock := range targetBlocks {
		alreadyExists := false
		for _, size := range targetSizes {
			if size.Equals(targetBlock.Size) {
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			targetSizes = append(targetSizes, targetBlock.Size)
		}
	}

	subBlocks := make(SubBlocks, 0)
	for _, targetSize := range targetSizes {
		h.Logger.Debug(fmt.Sprintf("Target size: %d x %d", targetSize.X, targetSize.Y))
		subBlock, err := b.getBlocksFromBlock(h, targetSize)
		if err != nil {
			h.Logger.Error(fmt.Sprintf("Error getting blocks from block: %s", err))
			return 0, err
		}
		subBlocks[targetSize] = subBlock
	}

	count := 0
	for _, subBlock := range subBlocks {
		for _, block := range subBlock {
			match, err := block.doBlocksMatchAny(h, targetBlocks)
			if err != nil {
				h.Logger.Error(fmt.Sprintf("Error checking if blocks match: %s", err))
				return 0, err
			}
			count += match
		}
	}
	return count, nil
}

// CountBlocks returns the number of times a Block appears in the puzzle
func (p *Puzzle) CountBlocksSameSize(h *common.Helpers, targets []Sets, rotate bool) (int, error) {
	if len(targets) == 0 {
		h.Logger.Error("No targets to count")
		return 0, fmt.Errorf("No targets to count")
	}
	for _, target := range targets {
		h.Logger.Debug(fmt.Sprintf("Target: %d x %d", len(target), len(target[0])))
		if len(targets[0]) == 0 {
			h.Logger.Error("No target size Y")
			return 0, fmt.Errorf("No target size Y")
		}
		if len(targets[0][0]) == 0 {
			h.Logger.Error("No target size X")
			return 0, fmt.Errorf("No target size X")
		}
	}
	return p.countBlockInBlockSameSize(h, targets, true)
}
