package txpool

import (
	"fmt"
	"sync"
)

type TxPool struct {
	mu   sync.Mutex
	Rows map[int][]string
}

func NewTxPool() *TxPool {
	return &TxPool{Rows: make(map[int][]string)}
}

func (tp *TxPool) AddTx(row int, tx string) {
	tp.mu.Lock()
	defer tp.mu.Unlock()
	tp.Rows[row] = append(tp.Rows[row], tx)
}

func (tp *TxPool) Process() {
	var wg sync.WaitGroup
	for row, txs := range tp.Rows {
		wg.Add(1)
		go func(r int, t []string) {
			defer wg.Done()
			fmt.Printf("Processing Row %d: %v\n", r, t)
		}(row, txs)
	}
	wg.Wait()
}
type Transaction struct {
	ID   string
	Data string
}

type TxPool struct {
	Transactions []Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{}
}

func (p *TxPool) AddTransaction(tx Transaction) {
	p.Transactions = append(p.Transactions, tx)
}

func (p *TxPool) GetTransactions() []Transaction {
	txs := p.Transactions
	p.Transactions = nil // flush pool
	return txs
}
