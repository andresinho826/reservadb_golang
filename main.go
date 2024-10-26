package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"reservas.com/reservas/entidad"
	"reservas.com/reservas/mappaing"
)

func main() {
	err := doMain()
	if err != nil {
		os.Exit(1)
	}

}

func doMain() error {
	log.Println("Application Running....")
	err := runServer()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func runServer() error {
	http.HandleFunc("/reservas/consultas", consultarReservas)
	http.HandleFunc("/reservas/crearusuario", crearUsuario)
	return http.ListenAndServe(":8095", nil)

}

func consultarReservas(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		fmt.Fprint(w, "Cosultando reserva")
	} else {
		fmt.Fprint(w, "Metodo incorrecto")
	}

}

func crearUsuario(w http.ResponseWriter, req *http.Request) {
	user := &mappaing.UsuarioStruct{}

	for name, values := range req.Header {
		for _, val := range values {
			if name != "Content-Length" {
				w.Header().Add(name, val)
			}

		}

	}

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	reqB, _ := io.ReadAll((req.Body))
	err := json.Unmarshal(reqB, user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp := entidad.CreateUsers(user)
	w.Write(resp)

}