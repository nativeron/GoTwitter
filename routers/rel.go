package routers

import (
	"net/http"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/models"
)

func Rel(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id es obligatorio", 400)
		return
	}

	var t models.Rel
	t.UserID = IdUser
	t.UserRelID = ID

	status, err := bd.InsertRel(t)
	if err != nil {
		http.Error(w, "ocurrio un error al intentar insertar la relacion", 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro insertar la relacion", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
