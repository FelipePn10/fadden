package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestKeypairGeneratePrivateKey: Testa a geração de uma chave privada.
// A função GeneratePrivateKey() deve retornar uma chave privada válida.
func TestKeypairSignVerifySuccess(t *testing.T) {
	privKey := GeneratePrivateKey()
	PublicKey := privKey.PublicKey()

	msg := []byte("Hello, World!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(PublicKey, msg))
}

// TestKeypairSignVerifyFail: Testa a verificação de uma assinatura com chaves diferentes.
// A verificação deve falhar.
func TestKeypaiSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	PublicKey := privKey.PublicKey()
	msg := []byte("Hello, World!")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(PublicKey, []byte("Hello, World")))

}
