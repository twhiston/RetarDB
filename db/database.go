package db

import (
	"errors"
)

type RDataBase struct {
	backupFile string
	dataBase   map[string]string
}

func NewRDataBase(backupFile string) *RDataBase {
	r := new(RDataBase)

	r.backupFile = backupFile
	r.dataBase = make(map[string]string)

	return r
}

func (r *RDataBase) Write(key string, value string) {
	r.dataBase[key] = value
}

func (r *RDataBase) Read(key string) (string, error) {
	if r.Has(key) {
		return r.dataBase[key], nil
	}

	return "", createKeyNotFoundError(key)
}

func (r *RDataBase) Has(key string) bool {
	if _, exists := r.dataBase[key]; exists {
		return true
	}

	return false
}

func (r *RDataBase) Delete(key string) error {
	if r.Has(key) {
		delete(r.dataBase, key)
		return nil
	}

	return createKeyNotFoundError(key)
}

func (r *RDataBase) Count() int {
	return len(r.dataBase)
}

func (r *RDataBase) Clear() {
	newDataBase := make(map[string]string)
	r.dataBase = newDataBase
}

func createKeyNotFoundError(key string) error {
	return errors.New(key + "is currently not stored inside the database")
}
