package migrate

import (
	"ETaalim/internals/model"
	"ETaalim/pkg/core"
	"fmt"
	"log"
)

func MakeMigrations() error {
	db := core.GetDBInstance()
	if err := db.AutoMigrate(&model.User{}, &model.Grade{}, &model.Teacher{}, &model.Student{}, &model.Class{}); err != nil {
		log.Fatal("\n[!] Error occurred during migration >>>", err)
		return fmt.Errorf("\nGORM_MIGRATION_ERROR: %+v.\n", err)
	}
	fmt.Println("[***] DB schema migrated successfully !")
	return nil
}
