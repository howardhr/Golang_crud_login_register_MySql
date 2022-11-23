package models

type Usuario struct {
	ID         int    `json:"id"`
	Documento  string `json:"documento"`
	Contrasena string `json:"contrasena"`
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Direccion  string `json:"direccion"`
	Telefono   int    `json:"telefono"`
}
