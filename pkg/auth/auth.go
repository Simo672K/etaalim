package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	AuthData AuthTokenData
	JWTToken
}

type AuthTokenData struct {
	FullName string `json:"fullName"`
	UniqueID string `json:"uniqueId"`
	Role     string `json:"role"`
}

type JWTToken interface {
	GenerateToken() (string, error)
	ValidateToken(string) bool
	GetPaylodData(*jwt.Token) map[string]interface{}
}

func (t *JWT) SetTokenData(atd AuthTokenData) {
	t.AuthData = atd
}

func (t *JWT) GenerateToken(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"fullName": t.AuthData.FullName,
		"uniqueId": t.AuthData.UniqueID,
		"role":     t.AuthData.Role,
		"exp":      time.Now().Add(time.Second * 30).Unix(),
	})

	tokenString, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *JWT) GenerateRefreshToken(secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func (t *JWT) ValidateToken(tokenString string, secret []byte) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	payload := t.GetPaylodData(token)

	return payload, nil
}

func (t *JWT) GetPaylodData(token *jwt.Token) map[string]interface{} {
	claims := token.Claims.(jwt.MapClaims)
	payloadData := make(map[string]interface{})

	for key, val := range claims {
		payloadData[key] = val
	}

	return payloadData
}
