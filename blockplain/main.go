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

	// Get transactions from the pool
	txs := pool.GetTransactions()

	// Create Genesis block
	genesis := block.NewBlock(0, 0, []string{"Genesis"}, "Root", "proof_genesis", "", "")
	if consensus.ValidateProof(genesis.Proof) {
		p.AddBlock(0, 0, genesis)
	}

	// Create subsequent blocks using transaction data
	x, y := 1, 0 // Start at (1, 0) after the genesis block
	for len(txs) > 0 {
		// Group transactions for the block
		blockData := []string{}
		for i := 0; i < 2 && len(txs) > 0; i++ {
			blockData = append(blockData, txs[0].Data)
			txs = txs[1:]
		}

		// Create a new block with the transaction data
		newBlock := block.NewBlock(x, y, blockData, "TimeLayer", fmt.Sprintf("proof_%d", x), genesis.Hash, "")
		if consensus.ValidateProof(newBlock.Proof) {
			p.AddBlock(x, y, newBlock)
		}

		// Move to the next block position
		y++
		if y > 2 { // You can adjust this limit as needed
			x++
			y = 0
		}
	}

	// Print the 2D plane
	explorer.PrintPlane(p)
}
