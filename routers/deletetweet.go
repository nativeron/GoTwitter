package routers

import (
	"net/http"

	"github.com/nativeron/GoTwitter/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe haber id", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IdUser)
	if err != nil {
		http.Error(w, "ocurrio un error", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application&json")
	w.WriteHeader(http.StatusCreated)
}
