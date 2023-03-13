package main

import (
	"fmt"
	"git/twittor/bd"
	"git/twittor/handlers"
	"log"
)

func main() {
	fmt.Println("Inicio programa")

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()

}
