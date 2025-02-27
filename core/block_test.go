package core

import (
	"testing"
	"time"

	"github.com/FelipePn10/fadden/crypto"
	"github.com/FelipePn10/fadden/types"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Validator)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())
}

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}
	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privKey))

	return b
}
