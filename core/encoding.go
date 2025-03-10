package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

// Define interfaces para codificação e decodificação.
// Encoder: Interface que define um método para serializar um objeto.
// Decoder: Interface que define um método para desserializar um objeto.
// Ambas utilizam tipos genéricos (T any) para permitir reutilização.

type Encoder[T any] interface {
	Encode(T) error
}

type Decoder[T any] interface {
	Decode(T) error
}

type GobTxEncoder struct {
	w io.Writer
}

// Encode implements Encoder.
func (g *GobTxEncoder) Encode(T any) error {
	panic("unimplemented")
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	gob.Register(elliptic.P256())
	return &GobTxEncoder{
		w: w,
	}
}

func (e *GobTxDecoder) Decode(tx *Transaction) error {
	return gob.NewDecoder(e.r).Decode(tx)
}

type GobTxDecoder struct {
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {
	gob.Register(elliptic.P256())
	return &GobTxDecoder{
		r: r,
	}
}

func (e *GobTxDecoder) Decoder(tx *Transaction) error {
	return gob.NewDecoder(e.r).Decode(tx)
}
