package models

// Pedido representa una orden de compra
type Pedido struct {
	ID           int           `json:"id"`
	ClienteID    int           `json:"cliente_id"`
	Productos    []Producto    `json:"productos"`
	Total        float64       `json:"total"`
	Estado       string        `json:"estado"`
	Pago         Pago          `json:"pago"`
	Calificacion *Calificacion `json:"calificacion,omitempty"`
}

// ActualizarEstado cambia el estado del pedido
func (p *Pedido) ActualizarEstado(nuevoEstado string) {
	p.Estado = nuevoEstado
}
