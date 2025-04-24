package crosschain

import (
	"fmt"
)

func CrossChainTx(fromChain, toChain int, data string) {
	fmt.Printf("Cross-chain Tx: [%d] -> [%d]: %s\n", fromChain, toChain, data)
}
