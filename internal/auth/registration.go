package auth

import "golang.org/x/crypto/bcrypt"

type Registration struct {
}

func NewRegistration() *Registration {
	return &Registration{}
}

func (r Registration) Registrate(username string, password string) error {
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
