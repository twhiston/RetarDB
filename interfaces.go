package main

type ClientHandler interface {
	handleClient(interface{})
}
