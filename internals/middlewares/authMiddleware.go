package middlewares

import (
	"ETaalim/api/handlers"
	"ETaalim/pkg/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jwtEnv = handlers.JwtEnv

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authBearer := c.GetHeader("Authorization")
		if len(authBearer) == 0 {
			fmt.Println("Error in auth bearer!")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized access",
			})
			return
		}

		tokenString := authBearer[len("Bearer "):]
		jwtToken := auth.JWT{}

		pd, err := jwtToken.ValidateToken(tokenString, []byte(jwtEnv.AccessTokenSecret))

		if err != nil {
			fmt.Println("Error, failed to validate token!")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Unauthorized access",
			})
			return
		}

		c.Set("payload", pd)
		c.Next()
	}
}
