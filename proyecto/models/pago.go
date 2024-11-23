package models

import "time"

// Pago almacena detalles del pago de un pedido
type Pago struct {
	ID     int       `json:"id"`
	Monto  float64   `json:"monto"`
	Fecha  time.Time `json:"fecha"`
	Metodo string    `json:"metodo"`
	Estado string    `json:"estado"`
}

// Método para realizar un pago
func (p *Pago) RealizarPago() bool {
	// Simulación de pago exitoso
	p.Estado = "Completado"
	return true
}
