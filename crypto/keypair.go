package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/FelipePn10/fadden/types"
)

//----------------------------------------// //----------------------------------------// //----------------------------------------// //----------------------------------------//
//Como funciona o ECDSA - Elliptic Curve Digital Signature Algorithm:
//O ECDSA funciona através da geração de chaves pública e privada, que são usadas para assinar e verificar transações digitais. A chave privada é usada para assinar uma transação, enquanto a chave pública é usada para verificar a autenticidade da assinatura. Para entender melhor como o ECDSA funciona, é importante conhecer os seguintes conceitos: curvas elipticas, funções hash, assinaturas digitais e chaves pública e privada. Faça algumas pesquisas antes de continuar para melhor entendimento.
//----------------------------------------// //----------------------------------------// //----------------------------------------// //----------------------------------------//

// Estrutura que representa uma chave privada.
type PrivateKey struct {
	key *ecdsa.PrivateKey
}

// Estrutura que representa uma chave pública. E usada para verificar assinaturas e derivar endereços.
type PublicKey struct {
	key *ecdsa.PublicKey
}

// --- Tanto o Private Key como o Public Key, são estruturas padrões em GO para representar chaves privadas e públicas de criptografia ECDSA. --- //

// Estrutura que representa uma assinatura ECDSA, contendo os valores r e s.
// Representa uma assinatura digital para validação.
type Signature struct {
	r, s *big.Int // big.Int: Representa um número inteiro grande. É usado para armazenar os valores r e s da assinatura.
}

//	Assina dados usando a chave privada ECDSA.
//
// ecdsa.Sign: Gera os componentes r e s da assinatura.
// rand.Reader: Garante que a geração da assinatura seja segura.
func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)
	if err != nil {
		return nil, err
	}

	return &Signature{
		r: r,
		s: s,
	}, nil
}

// Gera uma chave privada ECDSA usando a curva P-256 (secp256r1).
func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) // elliptic.P256(): Define a curva elíptica usada (NIST P-256). && rand.Reader: Gera números aleatórios seguros.
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

// Obtém a chave pública associada a uma chave privada.
// Retorna a chave pública associada a uma chave privada.
// A chave pública é derivada diretamente da chave privada no ECDSA.
// A chave pública é usada para verificar assinaturas e derivar endereços.
func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		key: &k.key.PublicKey,
	}
}

// Serializa a chave pública em um formato compacto (bytes).
// Retorna a chave pública serializada em um formato compacto.
// A chave pública é serializada em um formato compacto (33 bytes) que inclui o prefixo de compressão.
// O prefixo de compressão é um byte que indica se a chave pública é par ou ímpar.
func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.key, k.key.X, k.key.Y)
	// elliptic.MarshalCompressed: Converte as coordenadas (x, y) da chave pública em bytes compactos (ex: 0x02 ou 0x03 + coordenada x).
}

// Deriva um endereço (ex: de uma carteira) a partir da chave pública.
// Calcula o hash SHA-256 da chave pública serializada.
// Pega os últimos 28 bytes do hash.
// Converte esses bytes em um types.Address
func (k PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(h[len(h)-28:])
}

func (sig Signature) Verify(pubKey PublicKey, data []byte) bool { // Verifica se uma assinatura é válida para os dados e chave pública fornecidos.
	return ecdsa.Verify(pubKey.key, data, sig.r, sig.s) // ecdsa.Verify: Retorna true se a assinatura for válida.
}
