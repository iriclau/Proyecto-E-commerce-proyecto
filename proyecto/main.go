// Avance del Autonomo#2
//Nombres: Cristhian Sare, Freddy Campuzano, Iriany
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Ruta básica para probar el servidor
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Sistema de Gestión Iniciado")
	})

	// Crear ejemplos de uso
	http.HandleFunc("/ejemplo", ejemploHandler)

	log.Println("Servidor ejecutándose en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}

func ejemploHandler(w http.ResponseWriter, r *http.Request) {
	// Crear una instancia de Pedido con un Cliente y Productos
	cliente := models.Cliente{
		Usuario: models.Usuario{
			ID:     1,
			Nombre: "Juan Pérez",
			Correo: "juan@example.com",
			Rol:    "Cliente",
		},
	}

	producto := models.Producto{
		ID:     1,
		Nombre: "Laptop",
		Precio: 1200.50,
		Stock:  5,
	}

	carrito := models.Carrito{
		ID: 1,
	}
	carrito.AgregarProducto(producto)

	pedido := carrito.RealizarCompra(cliente.ID)
	calificacion := models.Calificacion{
		PedidoID:   pedido.ID,
		Puntuacion: 5,
		Comentario: "Excelente producto",
		Fecha:      time.Now(),
	}
	cliente.CalificarCompra(&pedido, calificacion)

	// Convertir pedido a JSON y enviar como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pedido)
}
