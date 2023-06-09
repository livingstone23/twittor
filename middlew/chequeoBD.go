package middlew

import (
	"git/twittor/bd"
	"net/http"
)

/*ChequeoBD: permite revisar la conexion a la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdidada con la Base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
