package models

// Categoria agrupa productos similares
type Categoria struct {
	ID        int        `json:"id"`
	Nombre    string     `json:"nombre"`
	Productos []Producto `json:"productos"`
}

// AgregarProducto añade un producto a la categoría
func (c *Categoria) AgregarProducto(producto Producto) {
	c.Productos = append(c.Productos, producto)
}

// EliminarProducto elimina un producto por su ID
func (c *Categoria) EliminarProducto(idProducto int) {
	for i, p := range c.Productos {
		if p.ID == idProducto {
			c.Productos = append(c.Productos[:i], c.Productos[i+1:]...)
			break
		}
	}
}
