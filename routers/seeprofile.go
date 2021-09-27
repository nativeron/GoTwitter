package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nativeron/GoTwitter/bd"
)

/*extrae valores de profile*/
func SeeProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Enviar parametro id", http.StatusBadRequest)
		return
	}
	profile, err := bd.SearchProfile(ID)
	if err != nil {
		http.Error(w, "ocurrio un error"+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
