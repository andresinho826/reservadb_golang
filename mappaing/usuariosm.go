package mappaing

type UsuarioStruct struct {
	Primer_Nombre            string `json:"nombre"`
	Segundo_Nombre           string `json:"segundoNombre"`
	Primer_Apellido          string `json:"primerApellido"`
	Segundo_Apellido         string `json:"segundoApellido"`
	Numero_Identificacion    string `json:"numeroIdentificacion"`
	Direccion                string `json:"direccion"`
	Nacionalidad             string `json:"nacionalidad"`
	Nick_Name                string `json:"nickName"`
	Fecha_De_Nacimiento      string `json:"fechaDeNacimiento"`
	Celular                  string `json:"celular"`
	Pais                     string `json:"pais"`
	Provincia_O_Departamento string `json:"provinciaODepartamento"`
	Email                    string `json:"email"`
}

type CredUsStruct struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
	Token    string `json:"token"`
	Attemps  string `json:"attemps"`
}
