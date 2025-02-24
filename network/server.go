// Esse arquivo contém toda a logica de um servidor de rede, que é responsável por receber mensagens de todos os transportes e processá-las. O servidor é composto por um canal rpcChan que recebe mensagens de todos os transportes, um canal quitCh para sinalizar a parada do servidor e um método Start que inicia o servidor e processa as mensagens recebidas. O método initTransports é responsável por inicializar os transportes e processar as mensagens recebidas por eles. O método Start é um loop que processa as mensagens recebidas do canal rpcChan, do canal quitCh e do ticker. O loop é interrompido quando uma mensagem é recebida do canal quitCh. O método initTransports é chamado no início do método Start para inicializar os transportes e processar as mensagens recebidas por eles. O método Start é um loop que processa as mensagens recebidas do canal rpcChan, do canal quitCh e do ticker. O loop é interrompido quando uma mensagem é recebida do canal quitCh
package network

import (
	"fmt"
	"time"
)

// ServerOpts define as opções do servidor.
// Transports é uma lista de transportes que o servidor pode usar.
// Ex: LocalTransport, RemoteTransport
type ServerOpts struct {
	Transports []Trasport
}

// Server é a estrutura que representa um servidor.
type Server struct {
	ServerOpts               // Opções do servidor
	rpcChan    chan RPC      // Canal central para receber mensagens de todos os transports
	quitCh     chan struct{} // Canal para sinalizar parada do servidor
}

// NewServer cria um novo servidor com as opções especificadas.
// Inicializa os canais rpcChan e quitCh.
func NewServer(opts ServerOpts) *Server {
	return &Server{ // Retorna um ponteiro para a estrutura Server
		ServerOpts: opts,                   // Inicializa as opções do servidor
		rpcChan:    make(chan RPC, 1024),   // Canal bufferizado para 1024 mensagens
		quitCh:     make(chan struct{}, 1), // Canal bufferizado para 1 mensagem
	}
}

func (s *Server) Start() {
	s.initTransports()                        // Inicializa os transportes
	ticker := time.NewTicker(1 * time.Second) // Cria um ticker que dispara a cada segundo

free: // Rótulo para o loop
	for {
		select { // Seleciona o primeiro canal que estiver pronto
		case rpc := <-s.rpcChan:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh: // Recebe uma mensagem do canal quitCh
			break free
		case <-ticker.C: // Tarefas periódicas (ex: logs)
			fmt.Println("so stuff every x seconds")
		}
	}
	fmt.Println("Server shutdown")
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
