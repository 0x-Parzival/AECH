package txpool

import (
	"fmt"
	"sync"
)

type Transaction struct {
	ID   string
	Data string
}

type TxPool struct {
	mu           sync.Mutex
	Transactions []Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{}
}

func (p *TxPool) AddTransaction(tx Transaction) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Transactions = append(p.Transactions, tx)
}

func (p *TxPool) GetTransactions() []Transaction {
	p.mu.Lock()
	defer p.mu.Unlock()
	txs := p.Transactions
	p.Transactions = nil // flush pool
	return txs
}

func (p *TxPool) Process() {
	var wg sync.WaitGroup
	p.mu.Lock()
	txs := p.Transactions
	p.mu.Unlock()

	for _, tx := range txs {
		wg.Add(1)
		go func(t Transaction) {
			defer wg.Done()
			fmt.Printf("Processing Tx %s: %s\n", t.ID, t.Data)
		}(tx)
	}
	wg.Wait()
}
