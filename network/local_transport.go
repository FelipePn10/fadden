package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr                     // Endereço deste nó (ex: "LOCAL", "REMOTE")
	consumuch chan RPC                    // Canal de recebimento de mensagens (buffer de 1024)
	lock      sync.RWMutex                // Mutex para sincronização concorrente
	peers     map[NetAddr]*LocalTransport // Mapa de peers conectados
}

func NewLocalTransport(addr NetAddr) Trasport {
	return &LocalTransport{
		addr:      addr,
		consumuch: make(chan RPC, 1024), // Canal bufferizado para 1024 mensagens
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Consume() <-chan RPC {
	return t.consumuch
}

func (t *LocalTransport) Connect(tr Trasport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport)

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s: could not send message to %s", t.addr, to)
	}

	peer.consumuch <- RPC{
		From:    t.addr,
		Payload: payload,
	}
	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
