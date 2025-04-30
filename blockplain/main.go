package main

import (
	"blockplain/block"
	"blockplain/plane"
	"blockplain/explorer"
	"blockplain/txpool"
	"blockplain/consensus"
	"blockplain/bridge"
)

func main() {
	// Create two planes for source and destination
	sourcePlane := plane.NewPlane()
	destPlane := plane.NewPlane()

	// Initialize the transaction pool
	pool := txpool.NewTxPool()

	// Add some transactions
	pool.AddTransaction(txpool.Transaction{ID: "tx1", Data: "Alice → Bob"})
	pool.AddTransaction(txpool.Transaction{ID: "tx2", Data: "Bob → Charlie"})

	// Get the transactions
	txs := pool.GetTransactions()
	data := []string{}
	for _, tx := range txs {
		data = append(data, tx.Data)
	}

	// Create Genesis block for source plane
	genesis := block.NewBlock(0, 0, []string{"Genesis"}, "Root", "proof_genesis", "", "")
	if consensus.ValidateProof(genesis.Proof) {
		sourcePlane.AddBlock(0, 0, genesis)
	}

	// Create next block for source plane using transaction data
	b1 := block.NewBlock(1, 0, data, "TimeLayer", "proof_1", genesis.Hash, "")
	if consensus.ValidateProof(b1.Proof) {
		sourcePlane.AddBlock(1, 0, b1)
	}

	// Create the bridge
	br := bridge.NewBridge(sourcePlane, destPlane)

	// Relaying transactions or blocks
	br.RelayTransaction(data, "TimeLayer", "proof_2")
	br.RelayBlock(b1)

	// Print source and destination planes
	fmt.Println("Source Plane:")
	explorer.PrintPlane(sourcePlane)

	fmt.Println("Destination Plane:")
	explorer.PrintPlane(destPlane)
}
