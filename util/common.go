package util

import (
	"encoding/json"
	"os"

	"reservas.com/reservas/mapping"
)

func GetDbStruct() (*mapping.ConnectionStruct, error) {
	cnnStruct := &mapping.ConnectionStruct{}
	pathf := "./resources/dbconfig.json"
	file, err := ReadFiles(pathf)
	if err != nil {
		return cnnStruct, err
	}
	err = json.Unmarshal(file, cnnStruct)
	if err != nil {
		return cnnStruct, err
	}
	return cnnStruct, nil
}

func ReadFiles(pathFile string) ([]byte, error) {
	return os.ReadFile(pathFile)
}
