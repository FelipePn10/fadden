package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/FelipePn10/fadden/crypto"

	"github.com/FelipePn10/fadden/types"
)

// Header: Estrutura que representa o cabeçalho de um bloco.
type Header struct {
	Version       uint32     // Versão do bloco
	Datahash      types.Hash // Hash dos dados do bloco
	PrevBlockHash types.Hash // Hash do bloco anterior
	Timestamp     uint64     // Timestamp do bloco
	Height        uint32     // Altura do bloco
}

// Block: Estrutura que representa um bloco.
type Block struct {
	*Header                        // Cabeçalho do bloco
	Transactions []Transaction     // Transações do bloco
	Validator    crypto.PublicKey  // Chave pública do validador
	Signature    *crypto.Signature // Assinatura do bloco
	hash         types.Hash        // Versão em cache do hash do cabeçalho
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{Header: h, Transactions: txx}
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.HeaderData())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.HeaderData()) {
		return fmt.Errorf("block has invalid signature")
	}

	return nil
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}

func (b *Block) HeaderData() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(b.Header); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
