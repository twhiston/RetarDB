package main

import (
	"fmt"
	"net"
	"testing"
)

func TestRDatabase_CRUD(t *testing.T) {
	database := NewRDataBase("data.json")

	database.Write("test", "test")

	if !database.Has("test") {
		t.Error("Database should have key test")
	}

	if database.Has("missingKey") {
		t.Error("Database should not have key missingKey")
	}

	data, err := database.Read("anotherMissingKey")

	if nil == err {
		t.Error("Reading a missing key should return an error")
	}
	if data != "" {
		t.Error("Missing key data should be empty")
	}

	database.Delete("test")

	if database.Has("test") {
		t.Error("Key should not exist anymore after deleting")
	}
}

func TestFunctional(t *testing.T) {
	config := createTestConfig()

	dataBase := NewRDataBase(config.BackupFile)
	clientHandler := NewRClientHandlerTCP(dataBase)
	server := NewRServer(config.ListenHost, clientHandler)

	go server.Run()

	client, err := net.Dial("tcp", config.ListenHost)

	if err != nil {
		t.Error("Can't connect to server")
	}

	fmt.Fprintf(client, `{"command": "write", "key": "value", "value": "test"}`)
}

func createTestConfig() *Config {
	c := new(Config)

	c.ListenHost = "127.0.0.1:8003"
	c.BackupFile = "backup.json"
	c.BackupRate = 30
	c.ServerType = "tcp"

	return c
}
