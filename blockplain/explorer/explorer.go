package explorer

import (
	"fmt"
	"blockplain/plane"
)

func PrintPlane(p *plane.Plane) {
	for x, row := range p.Grid {
		for y, block := range row {
			fmt.Printf("Block[%d][%d]: %s\n", x, y, block.Hash)
		}
	}
}
