package main

import (
	"time"

	"github.com/FelipePn10/fadden/network"
)

func main() {
	// Cria dois "transports" locais
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	// Conecta os transports entre si
	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	// Goroutine para enviar mensagens periodicamente
	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello, World!"))
			time.Sleep(1 * time.Second)
		}
	}()

	// Configuração do servidor
	opts := network.ServerOpts{
		Transports: []network.Trasport{trLocal},
	}

	// Inicializa e inicia o servidor
	s := network.NewServer(opts)
	s.Start()
}
