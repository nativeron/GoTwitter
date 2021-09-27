package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/nativeron/GoTwitter/bd"
)

func GetAllTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe mandar id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "debe mandar página", http.StatusBadRequest)
		return
	}

	pag, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "debe mandar página con valor mayor a 0", http.StatusBadRequest)
		return
	}

	page := int64(pag)
	resp, correct := bd.GetAllTweets(ID, page)
	if correct == false {
		http.Error(w, "error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
