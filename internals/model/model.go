package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UniqueID string `gorm:"unique;size:255"`
	FistName string
	LastName string
	Email    *string `gorm:"unique"`
	Birthday *time.Time
	Teacher  Teacher `gorm:"null;foreignKey:UserUID;references:ID"`
	Student  Student `gorm:"null;foreignKey:UserUID;references:ID"`
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
