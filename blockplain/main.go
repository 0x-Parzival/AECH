package main

import (
	"blockplain/block"
	"blockplain/plane"
	"blockplain/explorer"
	"blockplain/txpool"
	"blockplain/consensus"
	"fmt"
)

func main() {
	p := plane.NewPlane()
	pool := txpool.NewTxPool()

	// Add transactions to the pool
	pool.AddTransaction(txpool.Transaction{ID: "tx1", Data: "Alice → Bob"})
	pool.AddTransaction(txpool.Transaction{ID: "tx2", Data: "Bob → Charlie"})
	pool.AddTransaction(txpool.Transaction{ID: "tx3", Data: "Charlie → Dave"})
	pool.AddTransaction(txpool.Transaction{ID: "tx4", Data: "Dave → Eve"})
	pool.AddTransaction(txpool.Transaction{ID: "tx5", Data: "Eve → Frank"})
	pool.AddTransaction(txpool.Transaction{ID: "tx6", Data: "Frank → Grace"})

	// Get transactions from the pool
	txs := pool.GetTransactions()

	// Create Genesis block
	genesis := block.NewBlock(0, 0, []string{"Genesis"}, "Root", "proof_genesis", "", "")
	if consensus.ValidateProof(genesis.Proof) {
		p.AddBlock(0, 0, genesis)
	}

	// Coordinates for the next blocks
	x, y := 1, 0 // Start at (1, 0) after the genesis block

	// Create subsequent blocks using transaction data
	for len(txs) > 0 {
		// Group transactions for the block (grouping 2 transactions per block)
		blockData := []string{}
		for i := 0; i < 2 && len(txs) > 0; i++ {
			blockData = append(blockData, txs[0].Data)
			txs = txs[1:]
		}

		// Create a new block with the transaction data
		newBlock := block.NewBlock(x, y, blockData, "TimeLayer", fmt.Sprintf("proof_%d", x), genesis.Hash, "")

		// Validate the proof and add the block
		if consensus.ValidateProof(newBlock.Proof) {
			p.AddBlock(x, y, newBlock)
		}

		// Move to the next block position (move down in y, or if y exceeds limit, move to next row (x))
		y++
		if y > 2 { // Change the row after 3 blocks (adjust as needed)
			x++
			y = 0
		}
	}

	// Print the 2D plane after block additions
	explorer.PrintPlane(p)
}
