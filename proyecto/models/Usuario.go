package models

// Usuario representa un usuario genérico del sistema
type Usuario struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
	Rol        string `json:"rol"` // "Cliente", "Administrador"
}

// Cliente extiende Usuario
type Cliente struct {
	Usuario          // Composición
	Pedidos []Pedido `json:"pedidos"`
}

// Administrador extiende Usuario
type Administrador struct {
	Usuario // Composición
}
