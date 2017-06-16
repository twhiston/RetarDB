package db

import (
	"encoding/json"
	"errors"
)

const MESSAGEKEY_READ = "read"
const MESSAGEKEY_WRITE = "write"
const MESSAGEKEY_DELETE = "delete"
const MESSAGEKEY_HAS = "has"

type RMessageHandler struct {
	dataBase *RDataBase
}

func NewRMessageHandler(dataBase *RDataBase) *RMessageHandler {
	h := new(RMessageHandler)
	h.dataBase = dataBase
	return h
}

func (h *RMessageHandler) HandleMessage(rawMessage []byte) *Response {
	message, err := h.createMessage(rawMessage)

	if nil != err {
		return h.createErrorResponse(err)
	}

	return h.createMessageResponse(message)
}

func (h *RMessageHandler) createMessage(rawMessage []byte) (*Message, error) {
	m := new(Message)

	jsonError := json.Unmarshal(rawMessage, &m)
	if nil != jsonError {
		return nil, errors.New("can't parse message")
	}

	return m, nil
}

func (h *RMessageHandler) createErrorResponse(err error) *Response {
	m := new(Response)
	m.Error = err.Error()
	return m
}

func (h *RMessageHandler) createMessageResponse(message *Message) *Response {
	switch message.Command {

	case MESSAGEKEY_READ:
		return h.createAction(message)
	case MESSAGEKEY_WRITE:
		return h.writeAction(message)
	case MESSAGEKEY_HAS:
		return h.hasAction(message)
	case MESSAGEKEY_DELETE:
		return h.deleteAction(message)
	}

	r := new(Response)
	r.Error = "Invalid command: " + message.Command

	return r
}

func (h *RMessageHandler) createAction(msg *Message) *Response {
	r := new(Response)

	if h.dataBase.Has(msg.Key) {
		r.Message, _ = h.dataBase.Read(msg.Key)
		return r
	}

	r.Error = "Can not read: " + msg.Key
	return r
}

func (h *RMessageHandler) writeAction(msg *Message) *Response {
	r := new(Response)

	if len(msg.Key) > 0 {
		h.dataBase.Write(msg.Key, msg.Value)
		r.Message = "Success"
		return r
	}

	r.Error = "Invalid key for writing"
	return r
}

func (h *RMessageHandler) hasAction(msg *Message) *Response {
	r := new(Response)

	r.Message = "no"

	if h.dataBase.Has(msg.Key) {
		r.Message = "yes"
	}

	return r
}

func (h *RMessageHandler) deleteAction(msg *Message) *Response {
	r := new(Response)

	if h.dataBase.Has(msg.Key) {
		h.dataBase.Delete(msg.Key)
		r.Message = "Success"
		return r
	}

	r.Error = "Key does not exist: " + msg.Key

	return r
}
