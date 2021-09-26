package routers

import (
	"encoding/json"
	"net/http"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/models"
)

/*crea en la db el registro de usuario*/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "email está vacio", 400)
		return
	}
	if len(t.Password) < 8 {
		http.Error(w, "password debe ser igual o mayor a 8 caracteres", 400)
		return
	}

	_, find, _ := bd.CheckUserExist(t.Email)
	if find == true {
		http.Error(w, "ya existe usuario con ese email", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "ocurrió un error"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se logró insertar el registro"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
