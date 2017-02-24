package main

import "fmt"

const LEVEL_NOTICE = 100
const LEVEL_WARN = 50
const LEVEL_ERROR = 0

var CURR_LEVEL = 50

func Log(level int, msg string) {
	if level >= CURR_LEVEL {
		fmt.Println(msg)
	}
}

func Notice(msg string) {
	Log(LEVEL_NOTICE, "INFO: "+msg)
}

func Warn(msg string) {
	Log(LEVEL_WARN, "WARNING: "+msg)
}

func Error(msg string) {
	Log(LEVEL_ERROR, "ERROR: "+msg)
}
