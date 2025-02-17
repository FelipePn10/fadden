package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	trb.Connect(tra)
	tra.Connect(trb)
	assert.Equal(t, tra.peers[trb.addr], trb)
	assert.Equal(t, tra.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {

	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	trb.Connect(tra)
	tra.Connect(trb)

	msg := []byte("Hello, World!")
	assert.Nil(t, tra.SendMessage(trb.addr, msg))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.addr)
}
