package core

import (
	"testing"

	"github.com/FelipePn10/fadden/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.From)
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}
	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

// func TestTxEncodeDecode(t *testing.T) {
// 	tx := randomBlockWithSignature(t, 0, types.Hash{})
// 	buf := &bytes.Buffer{}
// 	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))

// 	txDecoded := new(Transaction)
// 	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
// 	assert.Equal(t, tx, txDecoded)
// }

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}
	assert.Nil(t, tx.Sign(privKey))

	return tx
}
