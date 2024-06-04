package auth

import (
	"ETaalim/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtEnv utils.EnvJwtData

func init() {
	jwtEnv.LoadEnv()
}

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
}

func (t *JWT) SetTokenData(atd AuthTokenData) {
	t.AuthData = atd
}

func (t *JWT) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"fullName": t.AuthData.FullName,
		"uniqueId": t.AuthData.UniqueID,
		"role":     t.AuthData.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtEnv.AccessTokenSecret))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (t *JWT) ValidateToken(token string) error {

	return nil
}
