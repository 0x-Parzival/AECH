package main

import (
	"blockplain/block"
	"blockplain/plane"
	"blockplain/explorer"
)

func main() {
	p := plane.NewPlane()

	genesis := &block.Block{
		X: 0, Y: 0,
		Data: []string{"Genesis"},
		Context: "Root",
		Proof: "proof_genesis",
	}
	p.AddBlock(0, 0, genesis)

	b1 := &block.Block{
		X: 1, Y: 0,
		Data: []string{"Tx1"},
		Context: "TimeLayer",
		Proof: "proof_1",
	}
	p.AddBlock(1, 0, b1)

	explorer.PrintPlane(p)
}
