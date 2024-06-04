package handlers

import (
	"ETaalim/pkg/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

	token, err := jwtToken.GenerateToken()
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
	auth := c.GetHeader("Authorization")

	fmt.Println(auth)
}
