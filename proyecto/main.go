// Avance del primer trabajo
//Nombres: Cristhian Sare, Freddy Campuzano
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Servicio de Gestión de E-commerce")
}

// Estructura de producto
type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

// Estructura de usuario
type Usuario struct {
	ID     int
	Nombre string
	Email  string
}

// Función para obtener productos
func getProducts() []Producto {
	return []Producto{
		{ID: 1, Nombre: "Laptop", Precio: 1000.00, Cantidad: 10},
		{ID: 2, Nombre: "Smartphone", Precio: 500.00, Cantidad: 20},
	}
}
