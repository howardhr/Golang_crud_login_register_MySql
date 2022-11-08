package models

type Persona struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
