package main

import (
	"log"

	"github.com/IsmaEstrella/Twitor/bd"
	"github.com/IsmaEstrella/Twitor/handlers"
)

// Realizar las importaciones de los paquetes
func main() {
	if bd.ChecarConexion() == 0 {
		log.Fatal("No hay conexion")
	}
	handlers.Manejadores()
}
