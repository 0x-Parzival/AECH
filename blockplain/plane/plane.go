package plane

import "blockplain/block"

type Plane struct {
	Grid map[int]map[int]*block.Block
}

func NewPlane() *Plane {
	return &Plane{
		Grid: make(map[int]map[int]*block.Block),
	}
}

func (p *Plane) AddBlock(x, y int, b *block.Block) {
	if _, exists := p.Grid[x]; !exists {
		p.Grid[x] = make(map[int]*block.Block)
	}
	prevX := p.GetBlock(x-1, y)
	prevY := p.GetBlock(x, y-1)

	if prevX != nil {
		b.PrevHashX = prevX.Hash
	}
	if prevY != nil {
		b.PrevHashY = prevY.Hash
	}

	b.Hash = b.ComputeHash()
	p.Grid[x][y] = b
}

func (p *Plane) GetBlock(x, y int) *block.Block {
	if col, exists := p.Grid[x]; exists {
		return col[y]
	}
	return nil
}
