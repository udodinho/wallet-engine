package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

func CompareHashPassword(password string, hashedPass []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), hashedPass)
	return err == nil
}
