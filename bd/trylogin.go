package bd

import (
	"github.com/nativeron/GoTwitter/models"
	"golang.org/x/crypto/bcrypt"
)

/**check de login a la db*/
func TryLogin(email string, password string) (models.User, bool) {
	usu, find, _ := CheckUserExist(email)
	if find == false {
		return usu, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
