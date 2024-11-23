package models

// Carrito almacena productos agregados por un cliente
type Carrito struct {
	ID        int        `json:"id"`
	Productos []Producto `json:"productos"`
	Total     float64    `json:"total"`
}

// Métodos para Carrito
func (c *Carrito) AgregarProducto(producto Producto) {
	c.Productos = append(c.Productos, producto)
	c.CalcularTotal()
}

func (c *Carrito) EliminarProducto(idProducto int) {
	for i, p := range c.Productos {
		if p.ID == idProducto {
			c.Productos = append(c.Productos[:i], c.Productos[i+1:]...)
			break
		}
	}
	c.CalcularTotal()
}

func (c *Carrito) CalcularTotal() {
	var total float64
	for _, p := range c.Productos {
		total += p.Precio
	}
	c.Total = total
}
