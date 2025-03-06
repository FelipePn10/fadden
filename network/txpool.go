package network

import (
	"sort"

	"github.com/FelipePn10/fadden/core"
	"github.com/FelipePn10/fadden/types"
)

type TxMapSorter struct {
	transactions []*core.Transaction
}

func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {
	txx := make([]*core.Transaction, len(txMap))

	i := 0
	for _, val := range txMap {
		txx[i] = val
		i++
	}
	s := &TxMapSorter{txx}

	sort.Sort(s)
	return s
}

func (s *TxMapSorter) Len() int { return len(s.transactions) }
func (s *TxMapSorter) Swap(i, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}
func (s *TxMapSorter) Less(i, j int) bool {
	return s.transactions[i].FirstSeen() < s.transactions[j].FirstSeen()
}

// Deefinimos um mapa onde armazena transações, onde a chave é um type.Hash e o valor é
// é um ponteiro para uma transação.
type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

// Esta função inicializa o pool de transações e retorna um novo.
// Ele cria um novo mapa vazio para transactions e retorna um ponteiro para a estrutura Txpool récem-criada.
func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Transactions() []*core.Transaction {
	s := NewTxMapSorter(p.transactions)
	return s.transactions
}

// Adiciona uma transação ao pool de transações
func (p *TxPool) Add(tx *core.Transaction) error {
	hash := tx.Hash(core.TxHasher{}) // Calcula o hash da transação. Isso gera um ID p/ a transação
	if p.Has(hash) {                 // Verifica se a transação já existe, se existir retorna nada
		return nil
	}
	p.transactions[hash] = tx // Se a transação não estiver no pool, ela é adicionada ao mapa transactions usando hash como chave
	return nil
}

// Verifica se uma transação com um determinado hash já existe no pool. ELe faz isso verificando se o hash está presente no mapa transactions, retornando true se estiver ou false.
func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}

// Retorna o número de transações atualmente no pool. (tamanho do mapa transactions)
func (p *TxPool) Len() int {
	return len(p.transactions)
}

// FLush limpa o pool de transações, removendo todas as transações.
// Ele faz isso criando um novo mapa vazio e atribuindo-o a transactions, efetivamente descartando todas as transações anteriores.
func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}
