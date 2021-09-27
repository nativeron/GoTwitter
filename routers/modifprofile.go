package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/models"
)

func ModifProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode((&t))
	if err != nil {
		http.Error(w, "datos incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.EditProfile(t, IdUser)
	if err != nil {
		http.Error(w, "ocurrio un error al modificar"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "ocurrio un error"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
