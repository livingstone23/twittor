package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"git/twittor/middlew"
	"git/twittor/routers"
)

/*Manejadores, seteoi puet y pongo a escuchar mi servidor*/
func Manejadores() {
	router := mux.NewRouter()


	router.HandleFunc("/registro",middlew.ChequeoBD(routers.Registro)).Methods("POST")







	//abrimos el puerto
	PORT := os.Getenv("PORT")
	if PORT == "" { 
	PORT = "8080" 
	}

	handlers := cors.AllowAll().Handler(router) 
	log.Fatal(http.ListenAndServe(":"+PORT, handlers))
}		