package routers

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nativeron/GoTwitter/bd"
	"github.com/nativeron/GoTwitter/models"
)

var IdUser string
var EmailUser string

/*process token p/ extraer sus valores*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {

	myKey := []byte("twittergo22")
	claims := &models.Claim{}

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, find, _ := bd.CheckUserExist(claims.Email)
		if find == true {
			EmailUser = claims.Email
			IdUser = claims.ID.Hex()
		}
		return claims, find, IdUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
