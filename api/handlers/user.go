package handlers

import (
	"ETaalim/internals/model"
	"ETaalim/pkg/auth"
	"ETaalim/pkg/core"
	"fmt"
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var db = core.GetDBInstance()

func GetUsers(c *gin.Context) {
	var users []model.User

	if err := db.Find(&users).Error; err != nil {
		fmt.Println("Error while fetching data")
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var User model.User
	User.UniqueID = uuid.NewString()
	
	if err := c.BindJSON(&User); err != nil {
		fmt.Println("failed to create user")
	}
	
	hashedPassword, err := auth.HashPassword(User.Password)
	User.Password = hashedPassword
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err,
		})
	}

	db.Create(&User)
	c.JSON(http.StatusCreated, User)
}
