package routes

import (
	"ETaalim/api/handlers"
	"ETaalim/internals/middlewares"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	r.GET("/student/all", middlewares.AuthMiddleware(), handlers.GetStudentsHandler)
	r.POST("/student/new", handlers.CreateStudent)
}
