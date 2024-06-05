package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type EnvJwtData struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
}

func (e *EnvJwtData) LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load environment variables, make sure you have .env file!")
	}
	e.AccessTokenSecret = os.Getenv("TOKEN_SECRET_KEY")
	e.RefreshTokenSecret = os.Getenv("REFRESH_SECRET_KEY")

	if e.AccessTokenSecret == "" || e.RefreshTokenSecret == "" {
		panic("there is no accesstoken/refreshtoken")
	}
}
