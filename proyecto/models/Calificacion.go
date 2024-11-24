package models

import "time"

// Calificacion representa una evaluación de un pedido
type Calificacion struct {
	ID         int       `json:"id"`
	PedidoID   int       `json:"pedido_id"`
	Puntuacion int       `json:"puntuacion"`
	Comentario string    `json:"comentario"`
	Fecha      time.Time `json:"fecha"`
}

// DarCalificacion asigna puntuación y comentario a una calificación
func (c *Calificacion) DarCalificacion(puntuacion int, comentario string) {
	c.Puntuacion = puntuacion
	c.Comentario = comentario
	c.Fecha = time.Now()
}
