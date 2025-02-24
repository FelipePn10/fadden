package core

import (
	"crypto/sha256"

	"github.com/FelipePn10/fadden/types"
)

// Com essa interface é possível criar hash de tudo o que quisermos desde que implemente o método Hash.
// Um Hasher é muito simples, é uma interface que pega um T, e vai fazer hash desse Tipo e retornar um Hash.
// Hasher: Interface que define o método Hash.
type Hasher[T any] interface {
	Hash(T) types.Hash // Método Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {
	h := sha256.Sum256(b.HeaderData())
	return types.Hash(h)
}
