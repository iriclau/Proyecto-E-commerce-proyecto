package models

// Pedido representa un pedido realizado por un cliente
type Pedido struct {
	ID           int           `json:"id"`
	Cliente      Cliente       `json:"cliente"`
	Productos    []Producto    `json:"productos"`
	Total        float64       `json:"total"`
	Estado       string        `json:"estado"` // "Pendiente", "Pagado", "Enviado", etc.
	Pago         Pago          `json:"pago"`
	Calificacion *Calificacion `json:"calificacion,omitempty"`
}

// Método para actualizar estado
func (p *Pedido) ActualizarEstado(nuevoEstado string) {
	p.Estado = nuevoEstado
}
