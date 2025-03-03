package core

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Blockchain struct { // Estrutura que representa a blockchain.
	store     Storage   // Armazenamento dos blocos
	headers   []*Header // Slice contendo os headers dos blocos da blockchain
	validator Validator // Validação dos blocos antes de serem adicionados
}

// Inicializa o store indicando que os blocos serão armazenados em memória,
// Define um validador de blocos,
// E adiciona o bloco gênisis (primeiro bloco da blockchain - o bloco gênisis é comum em todos os projetos blockhain espalhados pelo mundo - ).
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

// Primeiro o bloco passa pela validação via ValidateBlock, se for válido,
// o bloco é adicionado à blockchain
func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWiothoutValidation(b)
}

// Retorna o header do bloco na altura especificada.
// Se a altura informada for maior que a altura atual da blockchain, retorna erro
func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height (%d) too high", height)
	}
	return bc.headers[height], nil
}

// Retorna true se a blockchain possuir um bloco na altura fornecida
func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

// [0, 1, 2, 3] -> 4 len
// [0, 1, 2, 3] -> 3 height
// A altura da blockchain é o número total de blocos menos 1, pois os índices começam em zero.
func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

// Adiciona diretamente um bloco à blockchain sem validar.
// Registra no log a altura e o hash do bloco adicionado.
// Armazena o bloco no store.
func (bc *Blockchain) addBlockWiothoutValidation(b *Block) error {
	bc.headers = append(bc.headers, b.Header)
	logrus.WithFields(logrus.Fields{
		"height": b.Height,
		"hash":   b.Hash(BlockHasher{}),
	}).Info("adding new block")

	return bc.store.Put(b)
}
