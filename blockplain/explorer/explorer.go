package explorer

import (
	"fmt"
	"blockplain/plane"
	"blockplain/block"
)

// PrintPlane prints details of the entire 2D plane of blocks
func PrintPlane(p *plane.Plane) {
	for y := 0; y < len(p.Plane); y++ {
		for x := 0; x < len(p.Plane[y]); x++ {
			block := p.Plane[y][x]
			if block != nil {
				printBlockDetails(block)
			} else {
				fmt.Printf("No Block at [%d, %d]\n", x, y)
			}
		}
	}
}

// printBlockDetails prints the full details of a single block
func printBlockDetails(b *block.Block) {
	fmt.Printf("Block at [%d, %d] - Hash: %s\n", b.X, b.Y, b.Hash)
	fmt.Printf("  Previous Block Hash (PrevX, PrevY): %s, %s\n", b.PrevHashX, b.PrevHashY)
	fmt.Printf("  Data: %v\n", b.Data)
	fmt.Printf("  Context: %s\n", b.Context)
	fmt.Printf("  Proof: %s\n", b.Proof)
	fmt.Println()
}
