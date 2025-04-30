package crosschain

import (
	"fmt"
)

// CrossChainTx represents a transaction moving from one chain to another.
type CrossChainTx struct {
	FromChain int
	ToChain   int
	Data      string
}

// Send simulates sending a cross-chain transaction.
func (tx CrossChainTx) Send() {
	fmt.Printf("Cross-chain Tx: [%d] -> [%d]: %s\n", tx.FromChain, tx.ToChain, tx.Data)
}

// NewCrossChainTx creates a new cross-chain transaction instance.
func NewCrossChainTx(from, to int, data string) CrossChainTx {
	return CrossChainTx{
		FromChain: from,
		ToChain:   to,
		Data:      data,
	}
}
