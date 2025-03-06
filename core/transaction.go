package core

import (
	"fmt"

	"github.com/FelipePn10/fadden/crypto"
	"github.com/FelipePn10/fadden/types"
)

// Transaction: Estrutura que representa uma transação.
type Transaction struct {
	Data      []byte            // Dados da transação
	From      crypto.PublicKey  // Chave pública do remetente
	Signature *crypto.Signature // Guarda a assinatura digital da transação

	hash      types.Hash
	firstSeen int64
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Hash(h Hasher[*Transaction]) types.Hash {
	if tx.hash.IsZero() {
		tx.hash = h.Hash(tx)
	}
	return tx.hash
}

// Assina a transação com uma chave privada.
func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

// Valida se a transação foi assinada corretamente (se é legítima ou não).
func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}

func (tx *Transaction) SetFirstSeen(t int64) {
	tx.firstSeen = t
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.firstSeen
}
