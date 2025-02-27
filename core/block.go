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
	Version       uint32     // Identificador da versão do bloco
	Datahash      types.Hash // Hash do conteúdo das transações
	PrevBlockHash types.Hash // Hash do bloco anterior na rede
	Timestamp     uint64     // Marca o tempo de criação do bloco.
	Height        uint32     // Indica a posição do bloco na blockchain
}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)
	return buf.Bytes()
}

// Block: Estrutura que representa um bloco.
// O bloco contém um cabeçalho embutido (*Header), permitindo acesso direto aos seus campos:
// Transactions, Validator, Signature e Hash
type Block struct {
	*Header                        // Embeds Header (herança)
	Transactions []Transaction     // Lista de transações no bloco
	Validator    crypto.PublicKey  // Chave pública do validador do bloco
	Signature    *crypto.Signature // Assinatura do bloco
	hash         types.Hash        // Versão em cache do hash do cabeçalho
}

// Cria um novo bloco com um cabeçalho e uma lista de transações.
func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{Header: h, Transactions: txx}
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, *tx)
}

// Assina um Bloco - assina o cabeçalho do bloco usando uma chave privada
// privKey.Sign(b.HeaderData()): Gera a assinatura criptográfica
// b.Validator = privKey.PublicKey(): Armazena a chave pública do assinante
// b.Signature = sig: Salva a assinatura gerada
func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.Header.Bytes()) // Assina os dados do cabeçalho do bloco
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey() // Define o validator do bloco
	b.Signature = sig                 // Armazena a assinatura do bloco

	return nil
}

// Verifica a assinatura de um bloco. Garante que o bloco foi assinado corretamente, verifica
// a assinatura comparando com a chave pública do validador. Se não houver assinatura, ou a assinatura for inválida, um erro é retornado
func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	for _, tx := range b.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("block has invalid signature")
	}

	return nil
}

// Encode: Codifica um bloco em um escritor usando um codificador
func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

// Decode: Decodifica um bloco de um leitor usando um decodificador
func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

// Calcula o hash do bloco e armazena o resultado
// Se o b.Hash já tiver um valor, ele é retornado imediatamente
// Se for zero, usa um Hasher externo para calcular o hash do bloco
func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}
	return b.hash
}
