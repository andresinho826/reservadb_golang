package oracledb

import (
	"database/sql"
	"fmt"

	"reservas.com/reservas/util"

	_ "github.com/sijms/go-ora/v2"
)

func GgetDBConnection() (*sql.DB, error) {
	cnnStruct, err := util.GetDbStruct()

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
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
	//var query string
	//row := db.QueryRow("Select * from etc")
	//_ = row.Scan(&query)
	//fmt.Println("result", query)
}
