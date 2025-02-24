package core

import (
	"fmt"

	"github.com/FelipePn10/fadden/crypto"
)

// Transaction: Estrutura que representa uma transação.
type Transaction struct {
	Data []byte // Dados brutos da transação

	PublicKey crypto.PublicKey  // Chave pública do remetente
	Signature *crypto.Signature // Assinatura da transação
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}
	tx.PublicKey = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error {
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if tx.Signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
}
