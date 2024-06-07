package handlers

import (
	cred "ETaalim/internals/auth"
	"ETaalim/pkg/auth"
	"ETaalim/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var JwtEnv utils.EnvJwtData

func init() {
	JwtEnv.LoadEnv()
}

func AuthLoginHandler(c *gin.Context) {
	var credentials cred.Credentials

	var authTokenData *auth.AuthTokenData
	if err := c.Bind(&credentials); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	authTokenData, err := credentials.LoginWithCredentials()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	jwtToken := auth.JWT{}
	jwtToken.SetTokenData(*authTokenData)

	accessToken, err := jwtToken.GenerateToken([]byte(JwtEnv.AccessTokenSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	refreshToken, err := jwtToken.GenerateRefreshToken([]byte(JwtEnv.RefreshTokenSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

// Refreshing token handler
func AuthRefreshTokenHandler(c *gin.Context) {
	jwtToken := auth.JWT{}
	refreshTokenString := c.GetHeader("Authorization")[len("Bearer "):]

	if _, err := jwtToken.ValidateToken(refreshTokenString, []byte(JwtEnv.RefreshTokenSecret)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Expired refresh token",
		})
		return
	}

	accessToken, err := jwtToken.GenerateToken([]byte(JwtEnv.AccessTokenSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshTokenString,
	})
}
