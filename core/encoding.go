package core

import "io"

// Define interfaces para codificação e decodificação.
// Encoder: Interface que define um método para serializar um objeto.
// Decoder: Interface que define um método para desserializar um objeto.
// Ambas utilizam tipos genéricos (T any) para permitir reutilização.

type Encoder[T any] interface {
	Encode(io.Writer, T) error
}

type Decoder[T any] interface {
	Decode(io.Reader, T) error
}
