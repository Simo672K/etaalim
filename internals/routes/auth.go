package routes

import (
	"ETaalim/api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.POST("/api/auth/login", handlers.AuthLoginHandler)
}
