package middlew

import (
	"net/http"

	"github.com/nativeron/GoTwitter/bd"
)

/*middleware que me permite conocer el estado de la db*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "conexion perdida con la bd", 500)
			return
		}
		next.ServeHTTP(w, r)
	}

}
