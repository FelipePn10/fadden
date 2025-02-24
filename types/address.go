// Esse arquivo contém a definição do tipo Address, que é um array de 28 bytes. A função ToSlice converte o endereço para um slice de bytes. A função String converte o endereço para uma string hexadecimal. A função AddressFromBytes cria um endereço a partir de um slice de bytes.
// O arquivo também contém a definição de interfaces para codificação e decodificação. A interface Encoder define um método para serializar um objeto. A interface Decoder define um método para desserializar um objeto. Ambas utilizam tipos genéricos (T any) para permitir reutilização.

package types

import (
	"encoding/hex"
	"fmt"
)

// Address: Estrutura que representa um endereço.
// Um endereço é uma representação compacta de uma chave pública.
type Address [28]uint8

// ToSlice: Converte o endereço em um slice de bytes.
// Retorna um slice de bytes com 28 bytes.
func (a Address) ToSlice() []byte {
	b := make([]byte, 28)     // Cria um slice de bytes com 28 bytes.
	for i := 0; i < 28; i++ { // Copia os bytes do endereço para o slice.
		b[i] = a[i] // Copia o byte na posição i do endereço para a posição i do slice.
	}
	return b
}

// String: Converte o endereço em uma string hexadecimal.
// Retorna uma string hexadecimal com 56 caracteres.
func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

// AddressFromBytes: Converte um slice de bytes em um endereço.
// Deve receber um slice de bytes com 28 bytes.
// Retorna um endereço.
func AddressFromBytes(b []byte) Address {
	if len(b) != 28 { // Verifica se o slice de bytes tem 28 bytes.
		msg := fmt.Sprintf("given bytes with length %d should be 28", len(b))
		panic(msg) // Se não tiver 28 bytes, lança um erro.
	}

	var value [28]uint8       // Cria um array de 28 bytes.
	for i := 0; i < 28; i++ { // Copia os bytes do slice para o array.
		value[i] = b[i] // Copia o byte na posição i do slice para a posição i do array.
	}

	return Address(value) // Retorna o endereço.
}
