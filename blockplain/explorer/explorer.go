package explorer

import (
	"fmt"
	"blockplain/plane"
	"github.com/fatih/color" // For colored output
)

// PrintPlane will print the details of the entire 2D plane with better formatting
func PrintPlane(p *plane.Plane) {
	color.Green("------ Blockplain 2D Blockchain ------")
	for y := 0; y < len(p.Grid); y++ {
		fmt.Printf("\nRow %d:\n", y)
		for x := 0; x < len(p.Grid[y]); x++ {
			block := p.Grid[y][x]
			if block != nil {
				printBlockDetails(block)
			} else {
				color.Red("No Block at [%d, %d]\n", x, y)
			}
		}
	}
	fmt.Println("\n----------------------------------------")
}

// Helper function to print block details with better formatting
func printBlockDetails(b *block.Block) {
	color.Cyan("\nBlock at [%d, %d]", b.X, b.Y)
	color.Magenta("\nHash: %s", b.Hash)
	fmt.Printf("\n  Previous Block Hash (PrevX, PrevY): %s, %s", b.PrevHashX, b.PrevHashY)
	fmt.Printf("\n  Context: %s", b.Context)
	fmt.Printf("\n  Proof: %s", b.Proof)
	fmt.Printf("\n  Data: %v\n", b.Data)
}
