package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"git/twittor/bd"
	"git/twittor/models"
)

/*GraboTweet permite grabar el tweet en la base de datos*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	//Recibir el json en nuestro body y lo hemos copiado decodificado.
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar un registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
