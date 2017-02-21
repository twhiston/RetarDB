package main

const CONFIG_FILENAME = "retard.json"

func main() {
	config := CreateConfig(CONFIG_FILENAME)

	dataBase := NewRDataBase(config.BackupFile)
	backupHandler := NewRBackupHandler(dataBase, config.BackupRate)
	handler := NewClientHandler(dataBase)
	server := NewRServer(config.ListenHost, handler)

	go backupHandler.StartPeriodicBackup()

	server.Run()
}
