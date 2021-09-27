package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/jwt"
	"github.com/nativeron/GoTwitter/models"
)

/*realiza el login*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("context-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "usuario y/o contraseña invalida"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "mail es requerido", 400)
		return
	}
	doc, exist := bd.TryLogin(t.Email, t.Password)
	if exist == false {
		http.Error(w, "usuario y/o contraseña invalida", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "error al intentar general el token", 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	w.Header().Set("context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
