package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var fileavatar string = "uploads/avatars/" + IdUser + "." + extension

	f, err := os.OpenFile(fileavatar, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al subir la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error al copiar la imagen"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IdUser + "." + extension
	status, err = bd.EditProfile(user, IdUser)
	if err != nil || status == false {
		http.Error(w, "error al grabar avatar en la bd"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
