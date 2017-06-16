package main

import (
	"github.com/twhiston/RetarDB/db"
)

const CONFIG_FILENAME = "retard.json"

func main() {
	config := db.CreateConfig(CONFIG_FILENAME)

	dataBase := db.NewRDataBase(config.BackupFile)
	backupHandler := db.NewRBackupHandler(dataBase, config.BackupRate)
	handler := db.NewRClientHandlerTCP(dataBase)
	server := db.NewRServer(config.ListenHost, handler)

	go backupHandler.StartPeriodicBackup()

	server.Run()
}
