package models

// Usuario representa un usuario genérico del sistema.
type Usuario struct {
	ID         int    `json:"id"`         // ID único del usuario
	Nombre     string `json:"nombre"`     // Nombre del usuario
	Correo     string `json:"correo"`     // Correo electrónico
	Contrasena string `json:"contrasena"` // Contraseña para autenticación
	Rol        string `json:"rol"`        // Rol del usuario (ejemplo: "Cliente", "Administrador")
}

// Métodos para Usuario

// IniciarSesion valida si el correo y contraseña son correctos
func (u *Usuario) IniciarSesion(correo, contrasena string) bool {
	return u.Correo == correo && u.Contrasena == contrasena
}

// CerrarSesion simula el cierre de sesión
func (u *Usuario) CerrarSesion() bool {
	return true
}

// --------------------------------------------------------

// Cliente extiende Usuario con una lista de pedidos
type Cliente struct {
	Usuario          // Herencia de Usuario
	Pedidos []Pedido `json:"pedidos"` // Lista de pedidos realizados por el cliente
}

// Métodos para Cliente

// AgregarProductoACarrito permite a un cliente añadir productos al carrito
func (c *Cliente) AgregarProductoACarrito(carrito *Carrito, producto Producto) {
	carrito.Productos = append(carrito.Productos, producto)
	carrito.CalcularTotal() // Actualiza el total del carrito
}

// CalificarCompra permite a un cliente calificar un pedido
func (c *Cliente) CalificarCompra(pedido *Pedido, calificacion Calificacion) {
	pedido.Calificacion = &calificacion // Asigna la calificación al pedido
}

// --------------------------------------------------------

// Administrador extiende Usuario
type Administrador struct {
	Usuario // Herencia de Usuario
}

// Métodos para Administrador

// AgregarProducto agrega un nuevo producto a la lista de productos
func (a *Administrador) AgregarProducto(producto Producto, listaProductos *[]Producto) {
	*listaProductos = append(*listaProductos, producto)
}

// EliminarProducto elimina un producto de la lista por su ID
func (a *Administrador) EliminarProducto(idProducto int, listaProductos *[]Producto) {
	for i, p := range *listaProductos {
		if p.ID == idProducto {
			*listaProductos = append((*listaProductos)[:i], (*listaProductos)[i+1:]...)
			break
		}
	}
}

// ActualizarProducto actualiza un producto existente en la lista
func (a *Administrador) ActualizarProducto(idProducto int, productoActualizado Producto, listaProductos *[]Producto) {
	for i, p := range *listaProductos {
		if p.ID == idProducto {
			(*listaProductos)[i] = productoActualizado
			break
		}
	}
}
