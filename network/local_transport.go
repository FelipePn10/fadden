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

// Conecta dois transports para permitir comunicação.
// Usa sync.RWMutex para garantir acesso seguro ao mapa peers.
// Adiciona o peer ao mapa usando seu endereço como chave.
// tr.(*LocalTransport) força que todos os peers sejam do mesmo tipo (LocalTransport).
func (t *LocalTransport) Connect(tr Trasport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport) // Asserção de tipo para LocalTransport

	return nil
}

// Envia uma mensagem para um peer específico.
// Verifica se o peer existe no mapa.
// Se existir, envia um RPC para o canal consumuch do peer.
// A mensagem inclui o endereço do remetente (From) e o conteúdo (Payload).
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

// Expõe o canal de mensagens para que o servidor possa ler.
// O servidor lê do canal para processar mensagens recebidas.
func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
