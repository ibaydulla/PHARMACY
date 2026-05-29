package utils

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFrompassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPassword(password string, hashed string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))

}
