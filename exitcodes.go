package main

import (
	"fmt"
	"os"
)

const ERR_SERVER_START = 20
const ERR_LISTENER_ACCEPT = 21

const ERR_CONFIG_NOTFOUND = 30
const ERR_CONFIG_READ = 31
const ERR_CONFIG_UNMARSHAL = 32

func ErrorExit(msg string, errorCode int) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(errorCode)
}
