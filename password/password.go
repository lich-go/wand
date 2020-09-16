package password

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash :兼容PHP的非对称加密
func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verify :兼容PHP的非对称解密
func Verify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
