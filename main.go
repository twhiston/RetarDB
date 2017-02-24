package main

import ()

const CONFIG_FILENAME = "retard.json"

func main() {
	config := CreateConfig(CONFIG_FILENAME)

	dataBase := NewRDataBase(config.BackupFile)
	backupHandler := NewRBackupHandler(dataBase, config.BackupRate)
	handler := NewRClientHandlerTCP(dataBase)
	server := NewRServer(config.ListenHost, handler)

	go backupHandler.StartPeriodicBackup()

	server.Run()
}
