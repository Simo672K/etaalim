package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// GetStudentsHandler controller/handler for students route
func GetStudentsHandler(c *gin.Context) {
	students := []Student{
		{Name: "student 1", Age: 20},
		{Name: "student 2", Age: 20},
	}

	c.JSON(http.StatusOK, students)
}


func CreateStudent(c *gin.Context) {
}