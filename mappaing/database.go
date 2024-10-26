package mappaing

type ConnectionStruct struct {
	DbConfig struct {
		DB       string `json:"db"`
		Server   string `json:"server"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"dbConfig"`
}
