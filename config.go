package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func CreateConfig(fileName string) *Config {
	c := new(Config)

	fileContent := readFile(fileName)

	if len(fileContent) > 0 {
		jsonError := json.Unmarshal(fileContent, c)

		if nil != jsonError {
			ErrorExit("Error unmarshalling json: "+jsonError.Error(), ERR_CONFIG_UNMARSHAL)
		}
	} else {
		c.BackupRate = 30
		c.BackupFile = "backup.json"
		c.ListenHost = "127.0.0.1:8003"
		c.ServerType = "tcp"
	}

	return c
}

func readFile(fileName string) []byte {
	var content []byte

	_, err := os.Stat(fileName)
	if nil != err {
		Notice("No config file found. Using default settings")
	}

	content, readError := ioutil.ReadFile(fileName)

	if nil != readError {
		Error("Config file exists but is not readable. Using default settings")
	}

	return content
}
