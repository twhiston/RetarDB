package db

import (
	"encoding/json"
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

	data, err := database.Read("test")
	if nil != err {
		t.Error("Database should read key test")
	}
	if data != "test" {
		t.Error("Database key test should have value test")
	}

	if database.Has("missingKey") {
		t.Error("Database should not have key missingKey")
	}

	data, err = database.Read("anotherMissingKey")

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

var nestedExample = `
{
	"value1": "testing",
	"value2": {
		"this": "is",
		"a": "pretty classic",
		"json": "structure",
		"nesting": {
			"further": "should",
			"also": "work"
		}
	}
}
`

func TestNestedInsert(t *testing.T) {

	data := new(map[string]interface{})
	err := json.Unmarshal([]byte(nestedExample), data)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	database := NewRDataBase("data.json")
	database.WriteNested(*data)

	tv, err := database.Read("value1")
	if err != nil {
		t.Error(err.Error())
	} else {
		if tv != "testing" {
			t.Error("value1", "incorrect", "got:", tv, "expected", "testing")
		}
	}

	testvar := "value2.this"
	has := database.Has(testvar)
	if !has {
		t.Error("key", testvar, "should exist")
	}
	tv, err = database.Read(testvar)
	if err != nil {
		t.Error(err.Error())
	} else {
		if tv != "is" {
			t.Error("value2.this", "incorrect", "got:", tv, "expected", "is")
		}
	}

	testvar = "value2.json"
	has = database.Has(testvar)
	if !has {
		t.Error("key", testvar, "should exist")
	}
	tv, err = database.Read(testvar)
	if err != nil {
		t.Error(err.Error())
	} else {
		if tv != "structure" {
			t.Error("value2.this", "incorrect", "got:", tv, "expected", "structure")
		}
	}

	testvar = "value2.nesting.also"
	has = database.Has(testvar)
	if !has {
		t.Error("key", testvar, "should exist")
	}
	tv, err = database.Read(testvar)
	if err != nil {
		t.Error(err.Error())
	} else {
		if tv != "work" {
			t.Error("value2.this", "incorrect", "got:", tv, "expected", "work")
		}
	}

	testvar = "value2.nesting"
	has = database.Has(testvar)
	if !has {
		t.Error("key", testvar, "should exist")
	}
	tv, err = database.Read(testvar)
	if err == nil {
		t.Error("existing maps, that are not values should throw an error on read")
	}

	testvar = "value2.nesting.also.NOEXIST"
	has = database.Has(testvar)
	if has {
		t.Error("key", testvar, "should NOT exist")
	}
	tv, err = database.Read("value2.nesting.also.NOEXIST")
	if err == nil {
		t.Error("non existant keys should throw an error")
	}

}

func createTestConfig() *Config {
	c := new(Config)

	database := NewRDataBase("data.json")

	database.Write("test", "test")

	c.ListenHost = "127.0.0.1:8003"
	c.BackupFile = "backup.json"
	c.BackupRate = 30
	c.ServerType = "tcp"

	return c
}
