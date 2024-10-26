package entidad

import (
	"encoding/json"
	"fmt"

	"reservas.com/reservas/mappaing"
)

var myUser []*mappaing.UsuarioStruct

func CreateUsers(user *mappaing.UsuarioStruct) []byte {
	myUser = make([]*mappaing.UsuarioStruct, 0)
	myUser = append(myUser, user)

	for _, s := range myUser {
		res, _ := json.Marshal(s)
		fmt.Println(string(res))

	}
	reply, _ := json.Marshal(user)
	return reply

}
