package handlers

import (
	"ETaalim/pkg/auth"
	"ETaalim/pkg/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


var jwtEnv utils.EnvJwtData

func init() {
	jwtEnv.LoadEnv()
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

	token, err := jwtToken.GenerateToken([]byte(jwtEnv.AccessTokenSecret))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"accessToken": token,
	})
}

func AuthValidateTokenHandler(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	tokenString := authToken[len("Bearer "):]

	jwtToken := auth.JWT{}
	pd, err := jwtToken.ValidateToken(tokenString, []byte(jwtEnv.AccessTokenSecret))

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Authenticated successfully",
		"user":    pd,
	})
}
