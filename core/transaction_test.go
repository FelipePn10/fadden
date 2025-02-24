package core

// import (
// 	"testing"

// 	"github.com/FelipePn10/fadden/crypto"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSignTransaction(t *testing.T) {
// 	privKey := crypto.GeneratePrivateKey()
// 	tx := &Transaction{
// 		Data: []byte("foo"),
// 	}

// 	assert.Nil(t, tx.Sign(privKey))
// 	assert.NotNil(t, tx.PublicKey)
// }

// func TestVerifyTransaction(t *testing.T) {
// 	privKey := crypto.GeneratePrivateKey()
// 	tx := &Transaction{
// 		Data: []byte("foo"),
// 	}
// 	assert.Nil(t, tx.Sign(privKey))
// 	assert.Nil(t, tx.Verify())

// 	otherPrivKey := crypto.GeneratePrivateKey()
// 	tx.PublicKey = otherPrivKey.PublicKey()

// 	assert.NotNil(t, tx.Verify())
// }
