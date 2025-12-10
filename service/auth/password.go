package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func IsCorrectPassword(hashed, raw []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashed, raw)
	
	return err == nil
}