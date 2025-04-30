package plane

import "blockplain/block"

// Plane represents a 2D grid of blocks
type Plane struct {
	Grid map[int]map[int]*block.Block
}

// NewPlane creates a new empty 2D grid
func NewPlane() *Plane {
	return &Plane{
		Grid: make(map[int]map[int]*block.Block),
	}
}

// AddBlock inserts a block at position (x, y) in the grid
// It automatically sets PrevHashX and PrevHashY based on adjacent blocks
func (p *Plane) AddBlock(x, y int, b *block.Block) {
	if _, exists := p.Grid[x]; !exists {
		p.Grid[x] = make(map[int]*block.Block)
	}

	// Link to left (x-1) and above (y-1) blocks if they exist
	prevX := p.GetBlock(x-1, y)
	prevY := p.GetBlock(x, y-1)

	if prevX != nil {
		b.PrevHashX = prevX.Hash
	}
	if prevY != nil {
		b.PrevHashY = prevY.Hash
	}

	// Recompute hash with updated prev hashes
	b.Hash = b.ComputeHash()
	p.Grid[x][y] = b
}

// GetBlock retrieves the block at position (x, y) if it exists
func (p *Plane) GetBlock(x, y int) *block.Block {
	if col, exists := p.Grid[x]; exists {
		return col[y]
	}
	return nil
}
