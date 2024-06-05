package handlers

import (
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
	var atd auth.AuthTokenData
	if err := c.Bind(&atd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	jwtToken := auth.JWT{}
	jwtToken.SetTokenData(atd)

	at, err := jwtToken.GenerateToken([]byte(JwtEnv.AccessTokenSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	rt, err := jwtToken.GenerateRefreshToken([]byte(JwtEnv.RefreshTokenSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  at,
		"refreshToken": rt,
	})
}

// Refreshing token handler
func AuthRefreshTokenHandler(c *gin.Context) {
	jwtToken := auth.JWT{}
	rtString := c.GetHeader("Authorization")[len("Bearer "):]

	if _, err := jwtToken.ValidateToken(rtString, []byte(JwtEnv.RefreshTokenSecret)); err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Expired refresh token",
		})
		return
	}

	at, err := jwtToken.GenerateToken([]byte(JwtEnv.AccessTokenSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  at,
		"refreshToken": rtString,
	})
}
