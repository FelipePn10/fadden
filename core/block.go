package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/FelipePn10/fadden/types"
)

// Header: Estrutura que representa o cabeçalho de um bloco.
type Header struct {
	Version   uint32     // Versão do bloco
	PrevBlock types.Hash // Hash do bloco anterior
	Timestamp uint64     // Timestamp do bloco
	Height    uint32     // Altura do bloco
	Nonce     uint64     // Número aleatório
}

// EncodeBinary: Codifica o cabeçalho em binário.
func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	} // Escreve o valor de h.Version no escritor w em little-endian (ordem dos bytes) (4 bytes)
	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	} // Escreve o valor de h.PrevBlock no escritor w em little-endian (32 bytes)
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	} // Escreve o valor de h.Timestamp no escritor w em little-endian (8 bytes)
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	} // Escreve o valor de h.Height no escritor w em little-endian (4 bytes)
	return binary.Write(w, binary.LittleEndian, &h.Nonce) // Escreve o valor de h.Nonce no escritor w em little-endian (8 bytes)
}

// DecodeBinary: Decodifica o cabeçalho em binário.
func (h *Header) DecodeBinary(r io.Reader) error {
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Read(r, binary.LittleEndian, &h.Nonce)
}

// Block: Estrutura que representa um bloco.
type Block struct {
	Header                     // Cabeçalho do bloco
	Transactions []Transaction // Transações do bloco
	hash         types.Hash    // Versão em cache do hash do cabeçalho
}

// Hash: Retorna o hash do bloco.
// Se o hash ainda não foi calculado, calcula o hash e armazena em b.hash.
func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}     // Cria um buffer de bytes
	b.Header.EncodeBinary(buf) // Codifica o cabeçalho do bloco no buffer

	if b.hash.IsZero() { // Se o hash ainda não foi calculado
		b.hash = types.Hash(sha256.Sum256(buf.Bytes())) // Calcula o hash do buffer e armazena em b.hash
	}

	return b.hash
}

// EncodeBinary: Codifica o bloco em binário.
// Codifica o cabeçalho e as transações do bloco em binário.
func (b *Block) EncodeBinary(w io.Writer) error {
	if err := b.Header.EncodeBinary(w); err != nil { // Codifica o cabeçalho do bloco
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil { // Codifica cada transação do bloco
			return err
		}
	}
	return nil
}

// DecodeBinary: Decodifica o bloco em binário.
func (b *Block) DecodeBinary(r io.Reader) error {
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions { // Decodifica cada transação do bloco em binário e armazena em b.Transactions
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil // Retorna nil se não houver erros
}
