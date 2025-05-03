package main

import (
	"blockplain/block"
	"blockplain/consensus"
	"blockplain/explorer"
	"blockplain/plane"
	"blockplain/txpool"
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
	pool.AddTransaction(txpool.Transaction{ID: "tx7", Data: "Grace → Harry"})
	pool.AddTransaction(txpool.Transaction{ID: "tx8", Data: "Harry → Ivan"})

	// Get all transactions from the pool
	txs := pool.GetTransactions()

	// Create Genesis block at (0, 0)
	genesis := block.NewBlock(0, 0, []string{"Genesis"}, "Root", "proof_genesis", "", "")
	if consensus.ValidateProof(genesis.Proof) {
		p.AddBlock(0, 0, genesis)
	}

	// Start placing blocks from (1, 0)
	x, y := 1, 0
	maxCols := 3 // Number of columns in the grid before moving to next row

	// Create subsequent blocks, 2 txs per block
	for len(txs) > 0 {
		blockData := []string{}
		for i := 0; i < 2 && len(txs) > 0; i++ {
			blockData = append(blockData, txs[0].Data)
			txs = txs[1:]
		}

		newBlock := block.NewBlock(x, y, blockData, "TimeLayer", fmt.Sprintf("proof_%d_%d", x, y), genesis.Hash, "")

		if consensus.ValidateProof(newBlock.Proof) {
			fmt.Printf("Adding block at (%d, %d): %v\n", x, y, blockData)
			p.AddBlock(x, y, newBlock)
		}

		// Move to next position in grid (row-major)
		x++
		if x >= maxCols {
			x = 0
			y++
		}
	}

	// Print final 2D block grid
	explorer.PrintPlane(p)
}
