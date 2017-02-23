package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type SimpleClientHandler struct {
	dataBase       *RDataBase
	messageHandler *MessageHandler
}

func NewClientHandler(dataBase *RDataBase) *SimpleClientHandler {
	handler := new(SimpleClientHandler)

	handler.dataBase = dataBase
	handler.messageHandler = NewMessageHandler(dataBase)

	return handler
}

func (h *SimpleClientHandler) handleClient(conn net.Conn) {
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

func (h *SimpleClientHandler) readClientMessage(conn net.Conn) ([]byte, error) {
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
