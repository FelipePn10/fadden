package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Trasport
}

type Server struct {
	ServerOpts
	rpcChan chan RPC
	quitCh  chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcChan:    make(chan RPC, 1024),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(1 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcChan:
			fmt.Print("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("so stuff every x seconds")
		}
	}
	fmt.Println("Server shutdown")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Trasport) {
			for rpc := range tr.Consume() {
				s.rpcChan <- rpc
			}
		}(tr)
	}
}
