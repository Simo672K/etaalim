package model

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	AdminRole   Role = "admin"
	TeacherRole Role = "teacher"
	StudentRole Role = "student"
)

type User struct {
	gorm.Model
	UniqueID string     `json:"uniqueId" gorm:"unique;size:255"`
	FistName string     `json:"firstName"`
	LastName string     `json:"lastName"`
	Password string     `json:"password"`
	Email    *string    `json:"email" gorm:"unique"`
	Birthday *time.Time `json:"birthDay"`
	Role     Role       `json:"role" gorm:"type:enum('admin','teacher','student')"`
	Teacher  Teacher    `gorm:"null;foreignKey:UserUID;references:ID"`
	Student  Student    `gorm:"null;foreignKey:UserUID;references:ID"`
}

type Teacher struct {
	gorm.Model
	UserUID uint
}

type Grade struct {
	gorm.Model
	Title   string
	Classes []Class `gorm:"foreignKey:GradeId"`
}

type Student struct {
	gorm.Model
	UserUID uint
}

type Class struct {
	gorm.Model
	GradeId uint
}
