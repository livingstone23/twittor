package routers

import (
	"git/twittor/bd"
	"git/twittor/models"
	"io"
	"net/http"
	"os"
	"strings"
)

/*SubirBanner permite subir imagen de Avatar*/
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	//Abrimos el archivo, y le brindamos los permisos
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 666)
	if err != nil {
		http.Error(w, "Error al subir el Banner ! "+err.Error(), http.StatusBadRequest)
		return
	}

	//Copiamos el archivo renombrado
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el Banner en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
