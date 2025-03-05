package network

import (
	"fmt"
	"time"

	"github.com/FelipePn10/fadden/core"
	"github.com/FelipePn10/fadden/crypto"
	"github.com/sirupsen/logrus"
)

// ServerOpts define as opções do servidor.
// Transports é uma lista de transportes que o servidor pode usar.
// Ex: LocalTransport, RemoteTransport
type ServerOpts struct {
	Transports []Trasport
	BlockTime  time.Duration
	PrivateKey *crypto.PrivateKey
}

// Server é a estrutura que representa um servidor.
type Server struct {
	ServerOpts  // Opções do servidor
	blockTime   time.Duration
	memPool     *TxPool
	isValidator bool
	rpcChan     chan RPC      // Canal central para receber mensagens de todos os transports
	quitCh      chan struct{} // Canal para sinalizar parada do servidor
}

// NewServer cria um novo servidor com as opções especificadas.
// Inicializa os canais rpcChan e quitCh.
func NewServer(opts ServerOpts) *Server {
	return &Server{ // Retorna um ponteiro para a estrutura Server
		ServerOpts:  opts, // Inicializa as opções do servidor
		memPool:     NewTxPool(),
		blockTime:   opts.BlockTime,
		isValidator: opts.PrivateKey != nil,
		rpcChan:     make(chan RPC, 1024),   // Canal bufferizado para 1024 mensagens
		quitCh:      make(chan struct{}, 1), // Canal bufferizado para 1 mensagem
	}
}

func (s *Server) Start() {
	s.initTransports()                    // Inicializa os transportes
	ticker := time.NewTicker(s.blockTime) // Cria um ticker que dispara a cada segundo

free: // Rótulo para o loop
	for {
		select { // Seleciona o primeiro canal que estiver pronto
		case rpc := <-s.rpcChan:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh: // Recebe uma mensagem do canal quitCh
			break free
		case <-ticker.C: // Tarefas periódicas (ex: logs)
			if s.isValidator {
				s.createNewBlock()
			}
		}
	}
	fmt.Println("Server shutdown")
}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash) {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("transaction already in mempool")
		return nil
	}

	logrus.WithFields(logrus.Fields{
		"hash": hash,
	}).Info("adding new tx (%s) to the mempool")
	return s.memPool.Add(tx)
}

func (s *Server) createNewBlock() error {
	fmt.Println("creating a new block")
	return nil
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports { // Para cada transporte na lista de transportes
		go func(tr Trasport) {
			for rpc := range tr.Consume() { // Para cada mensagem recebida do canal de mensagens do transporte
				s.rpcChan <- rpc // Envia a mensagem para o canal rpcChan do servidor
			}
		}(tr)
	}
}
