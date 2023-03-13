package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

/*Manejadores, seteoi puet y pongo a escuchar mi servidor*/
func Manejadores() {
	router := mux.NewRouter()
	//abrimos el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" { 
	PORT = "8080" 
	}

	handlers := cors.AllowAll().Handler(router) 
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))
}		