package main

type Message struct {
	Command string `json:"command"`
	Key     string `json:"key"`
	Value   string `json:"value"`
}

type Response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type Config struct {
	BackupRate int    `json:"backup_rate"`
	BackupFile string `json:"backup_file`
	ListenHost string `json:"listen_host"`
}
