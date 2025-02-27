package core

import (
	"testing"

	"github.com/FelipePn10/fadden/types"
	"github.com/stretchr/testify/assert"
)

func TestAddBlock(t *testing.T) {
	bc := newBlockChainGenesis(t)
	lenBlocks := 1000

	for i := 0; i < lenBlocks; i++ {
		b := randomBlockWithSignature(t, uint32(i+1), getPrevblockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
	}

	assert.Equal(t, bc.Height(), uint32(lenBlocks))
	assert.Equal(t, len(bc.headers), lenBlocks+1)
	assert.NotNil(t, bc.AddBlock(randomBlock(89, types.Hash{})))
}

func TestNewBlockchain(t *testing.T) {
	bc := newBlockChainGenesis(t)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockChainGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestGetHeader(t *testing.T) {
	bc := newBlockChainGenesis(t)
	lenBlocks := 1000

	for i := 0; i < lenBlocks; i++ {
		b := randomBlockWithSignature(t, uint32(i+1), getPrevblockHash(t, bc, uint32(i+1)))
		assert.Nil(t, bc.AddBlock(b))
		header, err := bc.GetHeader(b.Height)
		assert.Nil(t, err)
		assert.Equal(t, header, b.Header)
	}
}

func TestAddBlockToHeigh(t *testing.T) {
	bc := newBlockChainGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 3, types.Hash{})))
}

func newBlockChainGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)
	return bc
}

func getPrevblockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}
