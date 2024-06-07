package auth

import (
	"ETaalim/pkg/auth"
	"ETaalim/internals/model"
	"ETaalim/pkg/core"
	"fmt"
)

var db = core.GetDBInstance()

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) LoginWithCredentials() (*auth.AuthTokenData ,error) {
	var User *model.User
	
	if c.Email == "" || c.Password == "" {
		return nil, fmt.Errorf("all fields are required")
	}
	
	if err := db.First(&User, "email = ?", c.Email).Error; err != nil {
		return nil, fmt.Errorf("email or user invalid")
	}
	
	match := auth.DoesPasswordMach(c.Password, User.Password)
	if !match {
		return nil, fmt.Errorf("email or user invalid")
	}
	
	userData := auth.AuthTokenData{
		FullName: User.FistName + " " + User.LastName,
		UniqueID: User.UniqueID,
		Role: string(User.Role),
	}
	
	
	return &userData, nil
}
