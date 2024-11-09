package entidad

import (
	"encoding/json"
	"fmt"

	"reservas.com/reservas/mapping"
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
