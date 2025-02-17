package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}

type Trasport interface {
	Consume() <-chan RPC
	Connect(Trasport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
