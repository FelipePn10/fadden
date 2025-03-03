package core

import "fmt"

// Define um contrato (interface) para qualquer validador de blocos.
// Qualquer estrutura que implemente essa interface deve possuir o método:
// ValidateBlock --> esse método recebe um bloco e retorna um erro se a validação falhar.
type Validator interface {
	ValidateBlock(*Block) error
}

// É a implementação concreta da interface Validator.
// Mantém uma referência (bc) para a blockchain associada.
type BlockValidator struct {
	bc *Blockchain
}

// Cria um novo validador de blocos associado a uma blockchain específica.
func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

// Esse é o coração da validação. Ele realiza várias verificações para garantir que o bloco seja válido antes de ser adicionado.
func (v *BlockValidator) ValidateBlock(b *Block) error {
	// Verifica se o bloco já existe
	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("chain already contains block (%d) with hash (%s)", b.Height, b.Hash(BlockHasher{}))
	}

	// Verifica se o bloco está na sequência correta.
	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("block (%s) too high", b.Hash(BlockHasher{}))
	}

	// Verfica se o hash do bloco anteriror é válido.
	prevHeader, err := v.bc.GetHeader(b.Height - 1)
	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	if hash != b.PrevBlockHash {
		return fmt.Errorf("the hash of the previous block (%s) is invalid", b.PrevBlockHash)
	}

	// Verifica a autenticidade do bloco
	if err := b.Verify(); err != nil {
		return err
	}

	return nil
}
