package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nativeron/GoTwitter/models"
)

/*genera encriptado con jwt*/
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("twittergo22")
	payload := jwt.MapClaims{
		"email":   t.Email,
		"name":    t.Name,
		"surname": t.Surname,
		"date":    t.Birth,
		"bio":     t.Bio,
		"ubi":     t.Ubi,
		"web":     t.Web,
		"_id":     t.ID.Hex(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
