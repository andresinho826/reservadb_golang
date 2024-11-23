package entidad

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"reservas.com/reservas/mapping"
	"reservas.com/reservas/oracledb"
)

var myUser []*mapping.UsuarioStruct

func CreateUsers(user *mapping.UsuarioStruct) []byte {
	myUser = make([]*mapping.UsuarioStruct, 0)
	myUser = append(myUser, user)

	for _, s := range myUser {
		res, _ := json.Marshal(s)
		fmt.Println(string(res))

	}
	reply, _ := json.Marshal(user)
	return reply

}

func FindUserById(id int) (*mapping.UsuarioStruct, error) {
	var user = &mapping.UsuarioStruct{}
	db, err := oracledb.GgetDBConnection()
	if err != nil {
		return user, err
	}
	qry := fmt.Sprintf("Select p_nombre, s_nombre from RES.Usuarios where id = %v", id)
	row := db.QueryRow(qry)
	defer db.Close()
	err = row.Scan(&user.Primer_Nombre, &user.Segundo_Nombre)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows were returned!")
		} else {
			fmt.Println("Error querying row:", err)
		}
		return user, err
	}
	//fmt.Println(rowData)
	return user, nil
}
