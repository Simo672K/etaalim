package routes

import (
	"ETaalim/api/handlers"
	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.RouterGroup) {
	r.GET("/students", handlers.GetStudentsHandler)
}
