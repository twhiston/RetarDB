package main

import (
	"net"
	"io"
	"fmt"
)

type PersistentClientHandler struct {
	dataBase *RDataBase
}

func NewClientHandler(dataBase *RDataBase) *PersistentClientHandler {
	handler := new(PersistentClientHandler)
	handler.dataBase = dataBase
	return handler
}

func (h *PersistentClientHandler) handleClient(conn net.Conn) {
	for {
		rawMessage, _ := h.readClientMessage(conn)
		fmt.Println(string(rawMessage))
	}
}

func (h *PersistentClientHandler) readClientMessage(conn net.Conn) ([]byte, error) {
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