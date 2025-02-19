package core

import "io"

// Transaction: Estrutura que representa uma transação.
type Transaction struct {
	Data []byte // Dados brutos da transação
}

// EncodeBinary: Codifica a transação em binário.
func (t *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}

// DecodeBinary Decodifica a transação em binário.
func (t *Transaction) DecodeBinary(r io.Reader) error {
	return nil
}
