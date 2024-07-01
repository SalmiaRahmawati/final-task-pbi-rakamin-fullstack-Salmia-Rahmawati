package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// func HashPass(p string) string {
// 	salt := 8
// 	password := []byte(p)
// 	hash, _ := bcrypt.GenerateFromPassword([]byte(password), salt)

// 	return string(hash)
// }

// func ComparePass(h, p []byte) bool {
// 	hash, pass := []byte(h), []byte(p)

// 	err := bcrypt.CompareHashAndPassword(hash, pass)
// 	return err == nil
// }
