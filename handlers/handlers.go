package handlers

import (
	"git/twittor/middlew"
	"git/twittor/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores, seteoi puet y pongo a escuchar mi servidor*/
func Manejadores() {
	router := mux.NewRouter()

	//Ruta para el registro de usuario
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	//abrimos el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handlers := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))
}
