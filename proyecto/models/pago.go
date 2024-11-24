package models

import "time"

// Pago contiene los detalles de una transacción
type Pago struct {
	ID     int       `json:"id"`
	Monto  float64   `json:"monto"`
	Fecha  time.Time `json:"fecha"`
	Metodo string    `json:"metodo"`
	Estado string    `json:"estado"`
}

// RealizarPago simula un pago exitoso
func (p *Pago) RealizarPago() bool {
	p.Estado = "Completado"
	p.Fecha = time.Now()
	return true
}
