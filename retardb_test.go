package main

import (
	"testing"
)

func TestDatabaseCRUD(t *testing.T) {
	database := NewRDataBase("data.json")

	database.Write("test", "test")

	if !database.Has("test") {
		t.Error("Database should have key test")
	}

	if database.Has("missingKey") {
		t.Error("Database should not have key missingKey")
	}

	data, err := database.Read("anotherMissingKey")

	if nil != err {
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

func TestReadJson(t *testing.T) {
	_ = CreateConfig(CONFIG_FILENAME)
}