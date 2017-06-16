package db

import (
	"errors"
	"log"
	"strings"
)

type dataBase map[string]interface{}

type RDataBase struct {
	backupFile       string
	dataBase         dataBase
	prependSeparator string
}

func NewRDataBase(backupFile string) *RDataBase {
	r := new(RDataBase)

	r.backupFile = backupFile
	r.dataBase = make(map[string]interface{})
	r.prependSeparator = "."

	return r
}

func (r *RDataBase) SetPrependSeparator(separator string) {
	r.prependSeparator = separator
}

func (r *RDataBase) Write(key string, value string) {
	r.dataBase[key] = value
}

func (r *RDataBase) Read(key string) (string, error) {

	//Explode the key by the delimiter
	stringSlice := strings.Split(key, r.prependSeparator)
	nextKey, nextKeys := r.keyShift(stringSlice)
	output, err := r.read(nextKey, nextKeys, r.dataBase)
	if err != nil {
		return "", err
	}

	return output, nil
}

// read recursive function that looks through a map[string]interface{} to try to resolve it to map[string]string types so a value can be returned
//TODO - currently if the last part of a tested key exists, but is a map and not a string endpoint it just returns an error, it would be better to have some more specific action so you could respond to it
func (r *RDataBase) read(key string, nextKeys []string, database dataBase) (string, error) {

	if r.has(key, nextKeys, database, false) {
		val := database[key]
		switch v := val.(type) {
		case string:

			if len(nextKeys) == 0 {
				//If there are no more keys left then we are good to go and can return
				return v, nil
			}

			index := len(nextKeys)
			if index > 0 {
				index -= 1
			}
			if key != nextKeys[index] {
				//It is possible to have a string result, but it not be far down the tree enough for the full key path.
				//In this case we must throw an error
				return "", createKeyNotFoundError(key)
			}
			return v, nil
		case map[string]interface{}:
			// array shift
			nextKey, nextKeys := r.keyShift(nextKeys)
			return r.read(nextKey, nextKeys, v)
		}
	}
	return "", createKeyNotFoundError(key)
}

func (r *RDataBase) keyShift(slice []string) (string, []string) {
	if len(slice) == 0 {
		log.Println("null key shift")
		return "", slice
	}
	return slice[0], slice[1:]
}

//WriteNested expects
func (r *RDataBase) WriteNested(values map[string]interface{}) {
	for k, val := range values {
		switch v := val.(type) {
		case string:
			r.dataBase[k] = v
		case map[string]interface{}:
			r.dataBase[k] = v
		default:
			log.Println("could not add values to database, unknown type, must be map[string]string or map[string]interface{}. All interface types must eventually resolve to map[string]string")
		}
	}
}

func (r *RDataBase) Has(key string) bool {
	stringSlice := strings.Split(key, r.prependSeparator)
	nextKey, nextKeys := r.keyShift(stringSlice)
	return r.has(nextKey, nextKeys, r.dataBase, true)
}

// Internal has representation, which allows for recursive calls when checking nested levels
func (r *RDataBase) has(key string, nextKeys []string, database dataBase, recurse bool) bool {

	if _, exists := database[key]; exists {
		switch v := database[key].(type) {
		case string:
			if len(nextKeys) == 0 {
				//If this is not the last key in the chain, but it is already a string value
				//then our intended value doesn't exist
				return true
			}
			return false
		case map[string]interface{}:
			if len(nextKeys) == 0 {
				//If we are at the final level of our key test, and something exists then we return true
				//This is debatable, as maybe we need to differentiate between exists and is key and exists and is map
				return true
			}
			//do more shifts and recursion
			nextKey, nextKeys := r.keyShift(nextKeys)
			return r.has(nextKey, nextKeys, v, recurse)
		}
	}
	return false
}

//TODO - refactor
func (r *RDataBase) Delete(key string) error {
	if r.Has(key) {
		delete(r.dataBase, key)
		return nil
	}

	return createKeyNotFoundError(key)
}

//TODO - refactor
func (r *RDataBase) Count() int {
	return r.count(r.dataBase)
}
func (r *RDataBase) count(base dataBase) int {
	return len(base)
}

func (r *RDataBase) Clear() {
	r.dataBase = make(dataBase)
}

func createKeyNotFoundError(key string) error {
	return errors.New("key: " + key + " not found in database")
}
