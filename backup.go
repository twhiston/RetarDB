package main

type RBackupHandler struct {
	rate int
	database *RDataBase
}

func NewRBackupHandler(database *RDataBase, rate int) *RBackupHandler {
	handler := new(RBackupHandler)
	handler.database = database
	handler.rate = rate
	return handler
}

func (h *RBackupHandler) StartPeriodicBackup() {

}