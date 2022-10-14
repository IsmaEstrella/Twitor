package middlew

import (
	"net/http"

	"github.com/IsmaEstrella/Twitor/bd"
)

func RevisionBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChecarConexion() == 0 {
			http.Error(w, "Conexion perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
