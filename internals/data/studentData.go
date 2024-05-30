package data

import (
	"ETaalim/pkg/core"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB

func init() {
	DBInstance = core.GetDBInstance()
}

func GetAllStudents() {

}

func GetStudentByID(id string) {

}
