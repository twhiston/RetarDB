package main

import (
	"net"
)

type RServer struct {
	host          string
	clientHandler ClientHandler
}

func NewRServer(host string, clientHandler ClientHandler) *RServer {
	s := new(RServer)

	s.host = host
	s.clientHandler = clientHandler

	return s
}

func (s *RServer) Run() {
	listener, serverStartError := net.Listen("tcp", s.host)
	if nil != serverStartError {
		ErrorExit("Error accpting from client: "+serverStartError.Error(), ERR_SERVER_START)
	}

	defer listener.Close()

	for {
		client, listenerAcceptError := listener.Accept()
		if nil != listenerAcceptError {
			ErrorExit("Error accepting client: "+listenerAcceptError.Error(), ERR_LISTENER_ACCEPT)
		}

		go s.handleClient(client)
	}
}

func (s *RServer) handleClient(client net.Conn) {
	go s.clientHandler.handleClient(client)
}
