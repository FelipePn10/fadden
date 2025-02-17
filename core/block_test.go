package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/FelipePn10/fadden/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode(t *testing.T) {
	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode := &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h.Version, hDecode.Version)

	var h2 Header
	if err := h2.DecodeBinary(buf); err != nil {
		t.Fatalf("DecodeBinary failed: %v", err)
	}

	if h.Version != h2.Version {
		t.Fatalf("Version mismatch: %d != %d", h.Version, h2.Version)
	}
	if h.PrevBlock != h2.PrevBlock {
		t.Fatalf("PrevBlock mismatch: %v != %v", h.PrevBlock, h2.PrevBlock)
	}
	if h.Timestamp != h2.Timestamp {
		t.Fatalf("Timestamp mismatch: %d != %d", h.Timestamp, h2.Timestamp)
	}
	if h.Height != h2.Height {
		t.Fatalf("Height mismatch: %d != %d", h.Height, h2.Height)
	}
	if h.Nonce != h2.Nonce {
		t.Fatalf("Nonce mismatch: %d != %d", h.Nonce, h2.Nonce)
	}
}
