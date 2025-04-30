package main

import (
	"blockplain/block"
	"blockplain/plane"
	"blockplain/explorer"
	"blockplain/txpool"
)

func main() {
	p := plane.NewPlane()
	pool := txpool.NewTxPool()

	// Add transactions to the pool
	pool.AddTransaction(txpool.Transaction{ID: "tx1", Data: "Alice → Bob"})
	pool.AddTransaction(txpool.Transaction{ID: "tx2", Data: "Bob → Charlie"})

	// Get transactions from the pool
	txs := pool.GetTransactions()
	data := []string{}
	for _, tx := range txs {
		data = append(data, tx.Data)
	}

	// Create Genesis block
	genesis := block.NewBlock(0, 0, []string{"Genesis"}, "Root", "proof_genesis", "", "")
	p.AddBlock(0, 0, genesis)

	// Create next block using tx data
	b1 := block.NewBlock(1, 0, data, "TimeLayer", "proof_1", genesis.Hash, "")
	p.AddBlock(1, 0, b1)

	// Print the 2D plane
	explorer.PrintPlane(p)
}
