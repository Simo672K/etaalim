package routes

import (
	"ETaalim/api/handlers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	r.GET("/student/all", handlers.GetStudentsHandler)
	r.POST("/student/new", handlers.CreateStudent)
}
