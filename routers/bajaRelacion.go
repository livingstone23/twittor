package routers

import (
	"git/twittor/bd"
	"git/twittor/models"
	"net/http"
)

/*BajaRelacion funcion que permite eliminar relacion*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioId = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al borrar la relación ", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
