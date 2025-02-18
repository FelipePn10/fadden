package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	trb.Connect(trb)
	tra.Connect(tra)
	// assert.Equal(t, trb.peers[trb.Addr()], trb)
	// assert.Equal(t, tra.peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello, World!")
	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.Addr())
}
