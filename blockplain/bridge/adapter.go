package bridge

import (
	"fmt"
	"blockplain/plane"
	"blockplain/block"
	"blockplain/consensus"
)

// Bridge struct
type Bridge struct {
	SourcePlane *plane.Plane
	DestPlane   *plane.Plane
}

// NewBridge creates a new Bridge
func NewBridge(source, dest *plane.Plane) *Bridge {
	return &Bridge{
		SourcePlane: source,
		DestPlane:   dest,
	}
}

// RelayTransaction relays a transaction from SourcePlane to DestPlane
func (b *Bridge) RelayTransaction(tx []string, context string, proof string) {
	// Create a new block with the transaction data
	block := block.NewBlock(0, 0, tx, context, proof, "", "")
	if consensus.ValidateProof(block.Proof) {
		// Add the block to the destination plane
		b.DestPlane.AddBlock(0, 0, block)
		fmt.Println("Transaction successfully relayed to destination plane!")
	} else {
		fmt.Println("Proof validation failed. Transaction not relayed.")
	}
}

// RelayBlock relays a full block (with its data and proof) to another plane
func (b *Bridge) RelayBlock(sourceBlock *block.Block) {
	// Validate the proof before adding to the destination plane
	if consensus.ValidateProof(sourceBlock.Proof) {
		// Add the block to the destination plane
		b.DestPlane.AddBlock(sourceBlock.X, sourceBlock.Y, sourceBlock)
		fmt.Printf("Block successfully relayed from [%d, %d] to [%d, %d]\n", sourceBlock.X, sourceBlock.Y, sourceBlock.X, sourceBlock.Y)
	} else {
		fmt.Println("Block proof validation failed. Block not relayed.")
	}
}

