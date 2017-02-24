package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type RClientHandlerTCP struct {
	dataBase       *RDataBase
	messageHandler *RMessageHandler
}

func NewRClientHandlerTCP(dataBase *RDataBase) *RClientHandlerTCP {
	handler := new(RClientHandlerTCP)

	handler.dataBase = dataBase
	handler.messageHandler = NewRMessageHandler(dataBase)

	return handler
}

func (h *RClientHandlerTCP) handleClient(client interface{}) {
	conn := client.(net.Conn)

	rawMessage, err := h.readClientMessage(conn)

	if nil != err {
		fmt.Println("close")
		conn.Close()
	}

	response := h.messageHandler.HandleMessage(rawMessage)
	jsonResponse, _ := json.Marshal(response)

	conn.Write(jsonResponse)
	conn.Close()
}

func (h *RClientHandlerTCP) readClientMessage(conn net.Conn) ([]byte, error) {
	tmp := make([]byte, 128)
	buf := make([]byte, 0, 2)

	for {
		n, err := conn.Read(tmp)
		if nil != err {
			if err != io.EOF {
				return buf, err
			}
			break
		}

		buf = append(buf, tmp[:n]...)
	}

	return buf, nil
}
