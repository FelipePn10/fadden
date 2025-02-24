// Esse arquivo é responsável por definir a interface de comunicação entre os nós da rede.
// A interface Transport define métodos para conectar nós, enviar mensagens e receber mensagens.
package network

type NetAddr string // Representa um endereço de rede. Ex: "LOCAL", "REMOTE" -> "192.168.1.1:3000"

// NetAddr: Endereço de um nó.
// RPC: Mensagem com remetente e conteúdo.
// Transport: Interface para comunicação entre nós.
type RPC struct {
	From    NetAddr // Quem enviou a mensagem (Endereço do remetente. Ex: "LOCAL", "REMOTE")
	Payload []byte  // Conteúdo da mensagem (dados brutos) (Dados enviados (ex: "Hello, World!", blocos de uma blockchain)
}

type Trasport interface {
	Consume() <-chan RPC               // Canal para receber mensagens do tipo RPC
	Connect(Trasport) error            // Conectar a outro Transport
	SendMessage(NetAddr, []byte) error // Enviar mensagem
	Addr() NetAddr                     // Retornar o endereço do Transport
}
