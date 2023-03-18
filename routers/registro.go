package routers

import (
	"encoding/json"
	"git/twittor/bd"
	"git/twittor/models"
	"net/http"
)

/*Registro: es la funcion para crear en la BD el regirstro de usuario*/
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La longitud del password es menor a 6", 400)
		return
	}

	//Validamos si el correo ya existe
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrodo con ese Email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	//Por si el estatus es vacio, da error pero no indica este
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario", 400)
		return
	}

	//si todo fue bien envio el status creado.
	w.WriteHeader(http.StatusCreated)

}
