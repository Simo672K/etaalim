package handlers

import (
	"ETaalim/internals/model"
	"ETaalim/pkg/core"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db = core.GetDBInstance()

func CreateUser(c *gin.Context) {
	var User model.User
	User.UniqueID = uuid.NewString()

	if err := c.BindJSON(&User); err != nil {
		fmt.Println("failed to create user")
	}

	db.Create(&User)
	c.JSON(http.StatusCreated, User)
}
