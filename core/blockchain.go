package core

import "fmt"

type Blockchain struct { // Estrutura que representa a blockchain.
	store     Storage
	headers   []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStorage(),
	}
	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWiothoutValidation(genesis)

	return bc, err
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWiothoutValidation(b)
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height (%d) too high", height)
	}
	return bc.headers[height], nil
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

// [0, 1, 2, 3] -> 4 len
// [0, 1, 2, 3] -> 3 height
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWiothoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)

	return bc.store.Put(b)
}
