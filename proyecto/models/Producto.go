package models

// Producto representa un artículo disponible para la venta
type Producto struct {
	ID          int      `json:"id"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion"`
	Precio      float64  `json:"precio"`
	Categorias  []string `json:"categorias"`
	Stock       int      `json:"stock"`
}

// ActualizarStock actualiza el inventario del producto
func (p *Producto) ActualizarStock(cantidad int) {
	p.Stock += cantidad
}
