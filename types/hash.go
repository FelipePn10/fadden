package types

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/exp/rand"
)

type Hash [32]uint8 // Definindo um tipo Hash que é um array de 32 bytes

// Verifica se o hash é zero. Se for zero, retorna true, senão, retorna false
func (h Hash) IsZero() bool {
	for i := 0; i < 32; i++ {
		if h[i] != 0 {
			return false
		}
	}
	return true
}

// Converte o hash para um slice de bytes
func (h Hash) ToSlice() []byte {
	b := make([]byte, 32)
	for i := 0; i < 32; i++ {
		b[i] = h[i]
	}
	return b
}

// Converte o hash para uma string. A string é o hash em hexadecimal. Ex: "a1b2c3d4.."
func (h Hash) String() string {
	return hex.EncodeToString(h.ToSlice())
}

// Função que cria um hash a partir de um slice de bytes. O slice de bytes deve ter 32 bytes, caso contrário, a função irá lançar um pânico.
func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("given bytes with length %d should be 32", len(b))
		panic(msg)
	}

	var value [32]uint8       // Cria um array de 32 bytes.
	for i := 0; i < 32; i++ { // Copia os bytes do slice para o array.
		value[i] = b[i] // Copia o byte na posição i do slice para a posição i do array.
	}

	return Hash(value) // Retorna o hash.
}

// Função que gera um slice de bytes aleatórios do tamanho especificado.
func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

// Função que gera um hash aleatório.
func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
