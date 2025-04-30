package main

import (
	"blockplain/block"
	"blockplain/plane"
	"blockplain/explorer"
	"blockplain/txpool"
	"blockplain/consensus"
)

func main() {
	p := plane.NewPlane()
	pool := txpool.NewTxPool()

	pool.AddTransaction(txpool.Transaction{ID: "tx1", Data: "Alice → Bob"})
	pool.AddTransaction(txpool.Transaction{ID: "tx2", Data: "Bob → Charlie"})

	txs := pool.GetTransactions()
	data := []string{}
	for _, tx := range txs {
		data = append(data, tx.Data)
	}

	genesis := block.NewBlock(0, 0, []string{"Genesis"}, "Root", "proof_genesis", "", "")
	if consensus.ValidateProof(genesis.Proof) {
		p.AddBlock(0, 0, genesis)
	}

	b1 := block.NewBlock(1, 0, data, "TimeLayer", "proof_1", genesis.Hash, "")
	if consensus.ValidateProof(b1.Proof) {
		p.AddBlock(1, 0, b1)
	}

	explorer.PrintPlane(p)
}
