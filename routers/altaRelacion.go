package routers

import (
	"git/twittor/bd"
	"git/twittor/models"
	"net/http"
)

/*AltaRelacion funcion que permite guardar la relacion*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro id", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioId = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar relacion ", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
