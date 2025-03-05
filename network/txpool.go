package network

import (
	"github.com/FelipePn10/fadden/core"
	"github.com/FelipePn10/fadden/types"
)

// Estrutura que armazena transações pendentes em um mapa
type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

// Função para inicializar o pool de transações
func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{})
	if p.Has(hash) {
		return nil
	}
	p.transactions[hash] = tx
	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}

func (p *TxPool) Len() int {
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}
