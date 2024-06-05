package routes

import (
	"ETaalim/api/handlers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine){
	r.GET("/user/all", handlers.GetUsers)
	r.POST("/user/new", handlers.CreateUser)
}