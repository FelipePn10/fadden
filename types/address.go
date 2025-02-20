package types

import (
	"encoding/hex"
	"fmt"
)

type Address [28]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 28)
	for i := 0; i < 28; i++ {
		b[i] = a[i]
	}
	return b
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 28 {
		msg := fmt.Sprintf("given bytes with length %d should be 28", len(b))
		panic(msg)
	}

	var value [28]uint8
	for i := 0; i < 28; i++ {
		value[i] = b[i]
	}

	return Address(value)
}
