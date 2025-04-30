package explorer

import (
	"fmt"
	"blockplain/block"
	"blockplain/plane"
	"github.com/fatih/color"
)

// PrintPlane will print the details of the entire 2D plane
func PrintPlane(p *plane.Plane) {
	color.Cyan("\n==================== Blockplain 2D Plane ====================")
	fmt.Println()
	for y := 0; y < len(p.Grid); y++ {
		fmt.Printf("Row %d:\n", y)
		for x := 0; x < len(p.Grid[y]); x++ {
			block := p.Grid[x][y]
			if block != nil {
				printBlockDetails(block)
			} else {
				color.Red("No Block at [%d, %d]\n", x, y)
			}
		}
		fmt.Println()
	}
	color.Cyan("==============================================================\n")
}

// Helper function to print block details
func printBlockDetails(b *block.Block) {
	color.Green("Block at [%d, %d]", b.X, b.Y)
	color.Magenta(" - Hash: %s", b.Hash)
	color.Yellow("  Previous Block Hash (PrevX, PrevY): %s, %s", b.PrevHashX, b.PrevHashY)
	color.White("  Data: %v", b.Data)
	color.Blue("  Context: %s", b.Context)
	color.HiWhite("  Proof: %s", b.Proof)
	color.Cyan("  Created at: %v", b.Timestamp)  // Assuming Timestamp field is added to the Block struct
	fmt.Println()
}
