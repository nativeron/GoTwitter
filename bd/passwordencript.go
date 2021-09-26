package bd

import "golang.org/x/crypto/bcrypt"

/*rutina que me permite encriptar password*/
func PasswordEncript(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}
