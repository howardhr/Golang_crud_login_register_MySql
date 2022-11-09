package models

type Usuario struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}
