package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	BackupRate int    `json:"backup_rate"`
	BackupFile string `json:"backup_file`
	ListenHost string `json:"listen_host"`
}

func CreateConfig(fileName string) *Config {
	c := new(Config)

	jsonError := json.Unmarshal(readFile(fileName), c)

	if nil != jsonError {
		ErrorExit("Error unmarshalling json: "+jsonError.Error(), ERR_CONFIG_UNMARSHAL)
	}

	return c
}

func readFile(fileName string) []byte {
	var content []byte

	_, err := os.Stat(fileName)
	if nil != err {
		ErrorExit("Error reading from file: "+fileName, ERR_CONFIG_NOTFOUND)
	}

	content, readError := ioutil.ReadFile(fileName)

	if nil != readError {
		ErrorExit("Error reading from file: "+fileName, ERR_CONFIG_READ)
	}

	return content
}
