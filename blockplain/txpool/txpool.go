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
	Transactions map[string]Transaction
}

// Constructor
func NewTxPool() *TxPool {
	return &TxPool{
		Transactions: make(map[string]Transaction),
	}
}

// Add transaction if not already present
func (p *TxPool) AddTransaction(tx Transaction) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, exists := p.Transactions[tx.ID]; !exists {
		p.Transactions[tx.ID] = tx
	}
}

// Return all transactions and flush pool
func (p *TxPool) GetTransactions() []Transaction {
	p.mu.Lock()
	defer p.mu.Unlock()

	txs := make([]Transaction, 0, len(p.Transactions))
	for _, tx := range p.Transactions {
		txs = append(txs, tx)
	}
	p.Transactions = make(map[string]Transaction) // flush
	return txs
}

// Concurrently process all transactions
func (p *TxPool) Process() {
	var wg sync.WaitGroup

	p.mu.Lock()
	txs := make([]Transaction, 0, len(p.Transactions))
	for _, tx := range p.Transactions {
		txs = append(txs, tx)
	}
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
