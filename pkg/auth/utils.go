package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func DoesPasswordMach(password, hashedPassword string)  bool{
	err:= bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
} 
