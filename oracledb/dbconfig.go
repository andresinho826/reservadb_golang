package oracledb

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
	"reservas.com/reservas/mapping"
	"reservas.com/reservas/util"
)

func GgetDBConnection() (*sql.DB, error) {
	cnnStruct, err := getDbStruct()

	if err != nil {
		return nil, err
	}
	var conectionString = fmt.Sprintf("oracle://%s:%s@%s:%s/%s",
		cnnStruct.DbConfig.User, cnnStruct.DbConfig.Password, cnnStruct.DbConfig.Server,
		cnnStruct.DbConfig.Port, cnnStruct.DbConfig.DB)
	db, err := sql.Open("oracle", conectionString)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDbStruct() (*mapping.ConnectionStruct, error) {
	cnnStruct := &mapping.ConnectionStruct{}
	pathf := "./resources/dbconfig.json"
	file, err := util.ReadFiles(pathf)
	if err != nil {
		return cnnStruct, err
	}
	err = json.Unmarshal(file, cnnStruct)
	if err != nil {
		return cnnStruct, err
	}
	return cnnStruct, nil
}
