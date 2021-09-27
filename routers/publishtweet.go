package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/models"
)

func PublishTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.InsertTweet{
		UserId:  IdUser,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(w, "ocurrio un error"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "no se logro insertar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
