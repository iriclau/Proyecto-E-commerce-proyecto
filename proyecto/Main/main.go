// Avance del Autonomo#3
//Nombres: Cristhian Sare, Freddy Campuzano, Iriany
/*Descripción: Avance pequeño donde podemos observar la estructura de los
productos y del usuario junto con la función getProduct*/
//Materia: Programación orientada a objetos.

package main

import (
	"Proyecto-E-commerce-Sare-Campuzano/models" // Importa tus modelos personalizados

	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	productos = []models.Producto{
		{ID: 1, Nombre: "Laptop", Precio: 1000.0, Stock: 10},
		{ID: 2, Nombre: "Smartphone", Precio: 700.0, Stock: 15},
		{ID: 3, Nombre: "Headphones", Precio: 100.0, Stock: 20},
	}
	carrito = models.Carrito{ID: 1}
)

func main() {
	r := gin.Default()

	// Cargar plantillas HTML
	r.LoadHTMLGlob("templates/*")

	// Servir archivos estáticos (CSS, JS)
	r.Static("/static", "./static")

	// Ruta para mostrar la página principal del e-commerce
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"productos":    productos,
			"totalCarrito": carrito.Total,
		})
	})

	// Ruta para agregar un producto al carrito
	r.POST("/add-to-cart", func(c *gin.Context) {
		var productoID int
		if err := c.ShouldBindJSON(&productoID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de producto inválido"})
			return
		}

		// Buscar producto por ID
		for _, p := range productos {
			if p.ID == productoID {
				carrito.AgregarProducto(p)
				c.JSON(http.StatusOK, gin.H{"message": "Producto agregado al carrito"})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
	})

	// Ruta para mostrar el contenido del carrito
	r.GET("/cart", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cart.html", gin.H{
			"productos": carrito.Productos,
			"total":     carrito.Total,
		})
	})

	// Iniciar el servidor
	r.Run(":8080")
}
